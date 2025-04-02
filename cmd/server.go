package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"

	"github.com/naka-gawa/grpc-benchtool/internal/metrics/datadog"
	"github.com/naka-gawa/grpc-benchtool/internal/server"
)

var port int

func newServerCmd() *cobra.Command {
	var cfg server.Config

	cmd := &cobra.Command{
		Use:   "server",
		Short: "Start gRPC benchmarking server",
		RunE: func(cmd *cobra.Command, args []string) error {
			var metricsClient *datadog.DatadogClient
			if cfg.DatadogCustomMetric {
				var err error
				metricsClient, err = datadog.NewDatadogClient()
				if err != nil {
					slog.Warn("failed to initialize Datadog client", slog.Any("error", err))
				}
			}

			strategy := &server.DefaultStrategy{
				ServerID:      cfg.ServerID,
				MetricsClient: metricsClient,
			}

			handler := server.NewBenchHandler(strategy)

			s, err := server.New(cfg, handler)
			if err != nil {
				slog.Error("failed to initialize server", slog.Any("error", err))
				return err
			}

			stopCh, ctx, cancel := handleSignals(cfg.Timeout)
			defer cancel()

			go func() {
				if err := s.Start(ctx); err != nil {
					slog.Error("server failed", slog.Any("error", err))
				}
			}()

			<-stopCh
			s.Shutdown(ctx)

			<-ctx.Done()
			return nil
		},
	}

	cmd.Flags().StringVar(&cfg.Host, "host", "0.0.0.0", "Host to bind the gRPC server")
	cmd.Flags().IntVar(&cfg.Port, "port", 50051, "Port to listen on")
	cmd.Flags().DurationVar(&cfg.Timeout, "timeout", 10*time.Second, "Shutdown timeout duration")
	cmd.Flags().StringVar(&cfg.ServerID, "server-id", "grpc-benchtool-server", "Optional server ID")
	cmd.Flags().BoolVar(&cfg.DatadogCustomMetric, "datadog-custom-metric", false, "Enable Datadog custom metric")

	return cmd
}

func handleSignals(timeout time.Duration) (chan os.Signal, context.Context, context.CancelFunc) {
	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	return stopCh, ctx, cancel
}
