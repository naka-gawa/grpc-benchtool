package main

import (
	"fmt"
	"log"
	"net"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	handler "github.com/naka-gawa/grpc-benchtool/internal/server"
	pb "github.com/naka-gawa/grpc-benchtool/proto/grpcbench"
)

var port int

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start gRPC benchmarking server",
	RunE: func(cmd *cobra.Command, args []string) error {
		addr := fmt.Sprintf(":%d", port)
		lis, err := net.Listen("tcp", addr)
		if err != nil {
			return fmt.Errorf("failed to listen: %w", err)
		}

		s := grpc.NewServer()
		reflection.Register(s)
		h := handler.NewBenchHandler("grpc-benchtool-server")
		pb.RegisterBenchServiceServer(s, h)

		log.Printf("server listening at %v\n", addr)
		return s.Serve(lis)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.Flags().IntVarP(&port, "port", "p", 50051, "Port to listen on")
}
