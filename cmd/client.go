package main

import (
	"fmt"
	"log/slog"

	"github.com/naka-gawa/grpc-benchtool/internal/client"
	"github.com/spf13/cobra"
)

var clientUsage = `grpc-benchtool client [flags]
Start gRPC benchmarking client.
This command allows you to start a gRPC benchmarking client that can send requests to a gRPC server.
You can specify the server address, request rate, duration of the test, and other parameters.
`

var clientExample = `# Start a gRPC benchmarking client with default settings
grpc-benchtool client

# Start a gRPC benchmarking client with unary mode
grpc-benchtool client --host localhost --port 50051 --mode unary --rps 10 --duration 10s --clientid client1

# Start a gRPC benchmarking client with streaming mode
grpc-benchtool client --host localhost --port 50051 --mode stream --rps 100 --duration 1m --clientid client2
`

func newClientCmd() *cobra.Command {
	cfg := &client.Config{}
	var profiler *client.Profiler
	cmd := &cobra.Command{
		Use:     "client",
		Short:   "Start gRPC benchmarking client",
		Long:    clientUsage,
		Example: clientExample,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := client.RunClient(*cfg); err != nil {
				return fmt.Errorf("failed to run client: %w", err)
			}
			return nil
		},
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			profiler = client.NewProfiler(*cfg)
			slog.Debug("config loaded", "enable_cpu_profile", cfg.EnableCPUProfile, "cpu_path", cfg.CpuProfilePath)
			if err := profiler.Start(); err != nil {
				fmt.Printf("failed to start CPU profile: %v\n", err)
			}
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			if err := profiler.Stop(); err != nil {
				fmt.Printf("failed to stop CPU profile: %v\n", err)
			}
		},
	}
	client.ClientFlags(cmd, cfg)
	return cmd
}
