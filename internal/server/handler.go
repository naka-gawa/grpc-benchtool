package server

import (
	"context"
	"io"
	"time"

	pb "github.com/naka-gawa/grpc-benchtool/proto/grpcbench"
)

type BenchHandler struct {
	pb.UnimplementedBenchServiceServer
	ServerID string
}

func NewBenchHandler(serverID string) *BenchHandler {
	return &BenchHandler{ServerID: serverID}
}

func (h *BenchHandler) UnaryTest(ctx context.Context, req *pb.TestRequest) (*pb.TestResponse, error) {
	now := time.Now().UnixNano()
	latency := now - req.SentUnixNano

	return &pb.TestResponse{
		ServerId:         h.ServerID,
		ReceivedUnixNano: now,
		LatencyNano:      latency,
	}, nil
}

func (h *BenchHandler) StreamTest(stream pb.BenchService_StreamTestServer) error {
	var (
		count    int64
		total    int64
		latSumMs float64
	)

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

	return stream.SendAndClose(&pb.StreamSummary{
		ServerId:         h.ServerID,
		ReceivedCount:    count,
		TotalBytes:       total,
		AverageLatencyMs: latSumMs / float64(count),
	})
}
