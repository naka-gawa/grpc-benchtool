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
	Role          string
	MetricsClient *datadog.DatadogClient
	ExtraTags     []string
}

func (s *DefaultStrategy) HandleUnary(ctx context.Context, req *pb.TestRequest) (*pb.TestResponse, error) {
	now := time.Now().UnixNano()
	latency := now - req.SentUnixNano

	if s.MetricsClient != nil {
		tags := []string{
			"type:unary",
			"server_id:" + s.ServerID,
		}
		tags = append(tags, s.ExtraTags...)
		_ = s.MetricsClient.SendGauge("grpc_benchtool.unary.latency", float64(latency), tags)
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
	var lateSumMs float64

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
		lateSumMs += latencyMs
	}

	if s.MetricsClient != nil && count > 0 {
		tags := []string{
			"type:stream",
			"server_id:" + s.ServerID,
		}
		tags = append(tags, s.ExtraTags...)
		_ = s.MetricsClient.SendGauge("grpc_benchtool.stream.latency", float64(lateSumMs), tags)
	}

	return stream.SendAndClose(&pb.StreamSummary{
		ServerId:      s.ServerID,
		ReceivedCount: count,
		TotalBytes:    total,
		LatencyMs:     lateSumMs,
	})
}
