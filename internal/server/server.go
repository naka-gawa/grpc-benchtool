package server

import (
	"context"
	"fmt"
	"log/slog"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/naka-gawa/grpc-benchtool/internal/interceptor"
	pb "github.com/naka-gawa/grpc-benchtool/proto/grpcbench"
)

type Server struct {
	cfg    Config
	grpc   *grpc.Server
	listen net.Listener
}

func New(cfg Config, handler pb.BenchServiceServer) (*Server, error) {
	addr := net.JoinHostPort(cfg.Host, fmt.Sprintf("%d", cfg.Port))
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("failed to listen on %s: %w", addr, err)
	}

	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptor.UnaryRequestIDInterceptor(),
			interceptor.UnaryLoggingInterceptor(),
		),
		grpc.ChainStreamInterceptor(
			interceptor.StreamRequestIDInterceptor(),
			interceptor.StreamLoggingInterceptor(),
		),
	)

	pb.RegisterBenchServiceServer(s, handler)

	reflection.Register(s)

	return &Server{
		cfg:    cfg,
		grpc:   s,
		listen: lis,
	}, nil
}

func (s *Server) Start(ctx context.Context) error {
	slog.Info("starting gRPC server", slog.String("addr", s.listen.Addr().String()))
	if err := s.grpc.Serve(s.listen); err != nil {
		return fmt.Errorf("failed to serve: %w", err)
	}
	return nil
}

func (s *Server) Shutdown(ctx context.Context) {
	slog.Info("shutting down gRPC server", slog.String("addr", s.listen.Addr().String()))

	done := make(chan struct{})
	go func() {
		s.grpc.GracefulStop()
		close(done)
	}()

	select {
	case <-ctx.Done():
		slog.Warn("shutdown timeout exceeded, forcing stop")
		s.grpc.Stop()
	case <-done:
		slog.Info("server stopped gracefully")
	}
}
