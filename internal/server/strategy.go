package server

import (
	"context"
	"io"
	"time"

	"github.com/naka-gawa/grpc-benchtool/internal/metrics/datadog"
	pb "github.com/naka-gawa/grpc-benchtool/proto/grpcbench"
)

type TestStrategy interface {
	HandleUnary(ctx context.Context, req *pb.TestRequest) (*pb.TestResponse, error)
	HandleStream(stream pb.BenchService_StreamTestServer) error
}

type DefaultStrategy struct {
	ServerID      string
	MetricsClient *datadog.DatadogClient
}

func (s *DefaultStrategy) HandleUnary(ctx context.Context, req *pb.TestRequest) (*pb.TestResponse, error) {
	now := time.Now().UnixNano()
	latency := now - req.SentUnixNano

	if s.MetricsClient != nil {
		_ = s.MetricsClient.SendGauge("grpc_benchtool.latency_ns", float64(latency), []string{
			"type:unary",
			"server_id:" + s.ServerID,
		})
	}

	return &pb.TestResponse{
		ServerId:         s.ServerID,
		ReceivedUnixNano: now,
		LatencyNano:      latency,
	}, nil
}

func (s *DefaultStrategy) HandleStream(stream pb.BenchService_StreamTestServer) error {
	var count int64
	var total int64
	var latSumMs float64

	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		count++
		total += int64(len(req.Payload))
		latencyMs := float64(time.Now().UnixNano()-req.SentUnixNano) / 1e6
		latSumMs += latencyMs
	}

	if s.MetricsClient != nil && count > 0 {
		_ = s.MetricsClient.SendGauge("grpc_benchtool.stream.avg_latency_ms", latSumMs/float64(count), []string{
			"type:stream",
			"server_id:" + s.ServerID,
		})
	}

	return stream.SendAndClose(&pb.StreamSummary{
		ServerId:         s.ServerID,
		ReceivedCount:    count,
		TotalBytes:       total,
		AverageLatencyMs: latSumMs / float64(count),
	})
}
