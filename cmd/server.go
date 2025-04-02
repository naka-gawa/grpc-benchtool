package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/naka-gawa/grpc-benchtool/internal/interceptor"
	"github.com/naka-gawa/grpc-benchtool/internal/server"
	pb "github.com/naka-gawa/grpc-benchtool/proto/grpcbench"
)

var port int

func newServerCmd() *cobra.Command {
	var (
		host     string
		port     int
		timeout  time.Duration
		serverID string
	)
	cmd := &cobra.Command{
		Use:   "server",
		Short: "Start gRPC benchmarking server",
		RunE: func(cmd *cobra.Command, args []string) error {
			addr := net.JoinHostPort(host, fmt.Sprintf("%d", port))
			lis, err := net.Listen("tcp", addr)
			if err != nil {
				slog.Error("failed to listen", slog.Any("error", err))
				return err
			}

			s, err := initializeServer(serverID)
			if err != nil {
				slog.Error("failed to initialize server", slog.Any("error", err))
				return err
			}

			stopCh, ctx, cancel := handleSignals(timeout)
			defer cancel()

			go startServer(s, lis, addr)

			<-stopCh
			shutdownServer(s, addr)

			<-ctx.Done()
			return nil
		},
	}

	cmd.Flags().StringVar(&host, "host", "0.0.0.0", "Host to bind the gRPC server")
	cmd.Flags().IntVar(&port, "port", 50051, "Port to listen on")
	cmd.Flags().DurationVar(&timeout, "timeout", 10*time.Second, "Shutdown timeout duration")
	cmd.Flags().StringVar(&serverID, "server-id", "grpc-benchtool-server", "Optional server ID")

	return cmd
}

func initializeServer(serverID string) (*grpc.Server, error) {
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

	reflection.Register(s)
	h := server.NewBenchHandler(serverID)
	pb.RegisterBenchServiceServer(s, h)

	return s, nil
}

func handleSignals(timeout time.Duration) (chan os.Signal, context.Context, context.CancelFunc) {
	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	return stopCh, ctx, cancel
}

func startServer(s *grpc.Server, lis net.Listener, addr string) {
	slog.Info("server listening", slog.String("address", addr))
	if err := s.Serve(lis); err != nil {
		slog.Error("failed to serve", slog.Any("error", err))
	}
}

func shutdownServer(s *grpc.Server, addr string) {
	slog.Info("shutting down gracefully...", slog.String("address", addr))
	s.GracefulStop()
}
