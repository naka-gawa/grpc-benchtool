package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/naka-gawa/grpc-benchtool/internal/logging"
	"github.com/spf13/cobra"
)

func main() {
	logging.Init()

	rootCmd := &cobra.Command{
		Use:   "grpc-benchtool",
		Short: "A gRPC benchmarking tool",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Use 'grpc-benchtool [command] --help'")
		},
	}

	rootCmd.AddCommand(newServerCmd())
	rootCmd.AddCommand(newClientCmd())

	if err := rootCmd.Execute(); err != nil {
		slog.Error("command execution failed", slog.Any("error", err))
		os.Exit(1)
	}
}
