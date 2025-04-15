package client

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/naka-gawa/grpc-benchtool/internal/metrics/datadog"
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
	EnableMetrics    bool
}

type Result struct {
	Success  bool
	Duration time.Duration
	ServerID string
	Error    error
}

// RunClient starts the gRPC benchmarking client. It connects to the gRPC server and sends requests based on the provided configuration.
func RunClient(cfg Config) error {
	client, conn, err := setupClient(cfg)
	if err != nil {
		return err
	}
	defer conn.Close()

	buffered, testCaseID := setupMetrics(cfg)
	if buffered != nil {
		buffered.Start()
		defer buffered.Stop()
	}

	resultCh := make(chan Result, 1000)
	go startResultCollector(resultCh)

	interval := time.Second / time.Duration(cfg.RPS)
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	timeout := time.After(cfg.Duration)
	wg := &sync.WaitGroup{}

LOOP:
	for {
		select {
		case <-timeout:
			slog.Info("test finished")
			break LOOP
		case <-ticker.C:
			wg.Add(1)
			go sendRequest(context.Background(), client, cfg, testCaseID, buffered, resultCh, wg)
		}
	}

	wg.Wait()
	close(resultCh)
	slog.Info("all requests done. finishing up")
	return nil
}

func setupClient(cfg Config) (pb.BenchServiceClient, *grpc.ClientConn, error) {
	addr := net.JoinHostPort(cfg.Host, cfg.Port)
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to server: %w", err)
	}

	client := pb.NewBenchServiceClient(conn)
	return client, conn, nil
}

func setupMetrics(cfg Config) (*datadog.BufferedClient, string) {
	testCaseID := uuid.New().String()
	slog.Info("generated test case id", "test_case_id", testCaseID)

	var buffered *datadog.BufferedClient
	if cfg.EnableMetrics {
		var err error
		buffered, err = datadog.NewBufferedClient(3 * time.Second)
		if err != nil {
			slog.Error("failed to initialize buffered datadog client", "error", err)
		}
	}

	return buffered, testCaseID
}

func startResultCollector(resultCh <-chan Result) {
	var total, success int
	var totalLatency time.Duration

	for res := range resultCh {
		total++
		if res.Success {
			success++
			totalLatency += res.Duration
		} else {
			slog.Warn("request failed", "error", res.Error)
		}
	}

	slog.Info("summary",
		"total", total,
		"success", success,
		"duration", totalLatency,
	)
}

func sendRequest(
	ctx context.Context,
	client pb.BenchServiceClient,
	cfg Config,
	testCaseID string,
	buffered *datadog.BufferedClient,
	resultCh chan<- Result,
	wg *sync.WaitGroup,
) {
	slog.Info("goroutine start")

	defer wg.Done()
	now := time.Now()
	req := &pb.TestRequest{
		ClientId:     cfg.ClientID,
		SentUnixNano: now.UnixNano(),
		PayloadBytes: 16,
		Payload:      []byte("hello gRPC!"),
	}

	timeoutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	resp, err := client.UnaryTest(timeoutCtx, req)
	if err != nil {
		if buffered != nil {
			tags := []string{
				fmt.Sprintf("client_id:%s", cfg.ClientID),
				fmt.Sprintf("mode:%s", cfg.Mode),
				fmt.Sprintf("test_case_id:%s", testCaseID),
			}
			buffered.Add("grpcbench.client.request.count", 1, tags)
			buffered.Add("grpcbench.client.failure.count", 1, tags)
		}

		resultCh <- Result{
			Success: false,
			Error:   err,
		}
		return
	}

	latency := time.Now().UnixNano() - req.SentUnixNano
	result := Result{
		Success:  true,
		ServerID: resp.ServerId,
		Duration: time.Duration(latency),
	}
	if buffered != nil {
		tags := []string{
			fmt.Sprintf("client_id:%s", cfg.ClientID),
			fmt.Sprintf("mode:%s", cfg.Mode),
			fmt.Sprintf("test_case_id:%s", testCaseID),
		}

		buffered.Add("grpcbench.client.request.count", 1, tags)

		if result.Success {
			buffered.Add("grpcbench.client.success.count", 1, tags)
			buffered.Add("grpcbench.client.latency.ms", result.Duration.Seconds()*1000, tags)
		} else {
			buffered.Add("grpcbench.client.failure.count", 1, tags)
		}
	}
	resultCh <- result
}
