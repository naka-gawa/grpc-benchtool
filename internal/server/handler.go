package server

import (
	"context"

	pb "github.com/naka-gawa/grpc-benchtool/proto/grpcbench"
)

type BenchHandler struct {
	pb.UnimplementedBenchServiceServer
	strategy TestStrategy
}

func NewBenchHandler(strategy TestStrategy) *BenchHandler {
	return &BenchHandler{strategy: strategy}
}

func (h *BenchHandler) UnaryTest(ctx context.Context, req *pb.TestRequest) (*pb.TestResponse, error) {
	return h.strategy.HandleUnary(ctx, req)
}

func (h *BenchHandler) StreamTest(stream pb.BenchService_StreamTestServer) error {
	return h.strategy.HandleStream(stream)
}
