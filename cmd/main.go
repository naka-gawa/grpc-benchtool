package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "grpc-benchtool",
	Short: "A gRPC benchmarking tool",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use 'grpc-benchtool [command] --help'")
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
