/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"

	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/lukehinds/grpc-auth/gen/go/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type apiServer struct {
	pb.UnimplementedAuthServer
}

func (s *apiServer) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	var token string
	fmt.Println(in)
	if in.Username == "luke" && in.Password == "password" {
		token = "1234567890"
	}
	return &pb.LoginResponse{
		Token: token,
	}, nil
}

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		lis, err := net.Listen("tcp", ":8080")
		if err != nil {
			log.Fatalln("Failed to listen:", err)
		}

		// Create a new gRPC server
		grpcServer := grpc.NewServer()

		pb.RegisterAuthServer(grpcServer, &apiServer{})

		// Serve gRPC Server
		log.Println("Serving gRPC on 0.0.0.0:8080")
		go func() {
			log.Fatalln(grpcServer.Serve(lis))
		}()

		conn, err := grpc.DialContext(
			context.Background(),
			"0.0.0.0:8080",
			grpc.WithBlock(),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			log.Fatalln("Failed to dial server:", err)
		}

		gwmux := runtime.NewServeMux()

		err = pb.RegisterAuthHandler(context.Background(), gwmux, conn)
		if err != nil {
			panic(err)
		}

		gwServer := &http.Server{
			Addr:    ":8090",
			Handler: gwmux,
		}

		log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
		log.Fatalln(gwServer.ListenAndServe())

	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
