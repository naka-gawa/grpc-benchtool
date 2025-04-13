package client

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"sync"
	"time"

	pb "github.com/naka-gawa/grpc-benchtool/proto/grpcbench"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Config struct {
	Host             string
	Port             string
	Mode             string
	RPS              uint16
	Duration         time.Duration
	ClientID         string
	EnableCPUProfile bool
	CpuProfilePath   string
}

type Result struct {
	Success  bool
	Latency  time.Duration
	ServerID string
	Error    error
}

// RunClient starts the gRPC benchmarking client. It connects to the gRPC server and sends requests based on the provided configuration.
func RunClient(cfg Config) error {
	addr := net.JoinHostPort(cfg.Host, cfg.Port)
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("failed to connect to server: %w", err)
	}
	defer conn.Close()

	client := pb.NewBenchServiceClient(conn)

	interval := time.Second / time.Duration(cfg.RPS)
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	timeout := time.After(cfg.Duration)
	count := 0

	wg := &sync.WaitGroup{}
	resultCh := make(chan Result, 1000)
	go func() {
		var total, success int
		var totalLatency time.Duration

		for res := range resultCh {
			total++
			if res.Success {
				success++
				totalLatency += res.Latency
			} else {
				slog.Warn("request failed", "error", res.Error)
			}
		}

		avg := time.Duration(0)
		if success > 0 {
			avg = totalLatency / time.Duration(success)
		}

		slog.Info("summary",
			"total", total,
			"success", success,
			"avg_latency", avg,
		)
	}()

LOOP:

	for {
		select {
		case <-timeout:
			slog.Info("test finished",
				slog.Int("total_requests", count),
				slog.String("duration", cfg.Duration.String()),
				slog.String("rps", fmt.Sprintf("%d", cfg.RPS)),
				slog.String("client_id", cfg.ClientID),
			)
			ticker.Stop()
			break LOOP

		case <-ticker.C:
			wg.Add(1)
			count++
			go func() {
				slog.Info("goroutine start")
				defer wg.Done()
				now := time.Now()
				req := &pb.TestRequest{
					ClientId:     cfg.ClientID,
					SentUnixNano: now.UnixNano(),
					PayloadBytes: 16,
					Payload:      []byte("hello gRPC!"),
				}

				ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
				defer cancel()

				resp, err := client.UnaryTest(ctx, req)
				if err != nil {
					resultCh <- Result{
						Success: false,
						Error:   err,
					}
					return
				}

				latency := time.Now().UnixNano() - req.SentUnixNano
				resultCh <- Result{
					Success:  true,
					ServerID: resp.ServerId,
					Latency:  time.Duration(latency),
				}
			}()

		}
	}
	wg.Wait()
	close(resultCh)
	slog.Info("all requests done. finishing up")
	return nil
}
