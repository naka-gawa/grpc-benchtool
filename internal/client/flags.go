package client

import (
	"time"

	"github.com/spf13/cobra"
)

// ClientFlags sets up the command-line flags for the gRPC benchmarking client.
func ClientFlags(cmd *cobra.Command, cfg *Config) {
	cmd.Flags().StringVar(&cfg.Host, "host", "localhost", "gRPC server host")
	cmd.Flags().StringVar(&cfg.Port, "port", "50051", "gRPC server port")
	cmd.Flags().StringVar(&cfg.Mode, "mode", "unary", "gRPC call mode (unary, stream)")
	cmd.Flags().Uint16Var(&cfg.RPS, "rps", 1, "Requests per second")
	cmd.Flags().DurationVar(&cfg.Duration, "duration", 10*time.Second, "Duration of the test (e.g., 10s, 1m)")
	cmd.Flags().StringVar(&cfg.ClientID, "clientid", "client", "Client ID for the test")
	cmd.Flags().BoolVar(&cfg.EnableCPUProfile, "enablecpuprofile", false, "Enable CPU profiling")
	cmd.Flags().StringVar(&cfg.CpuProfilePath, "cpuprofpath", "./bench/cpu.pprof", "Path to save CPU profile file")
}
