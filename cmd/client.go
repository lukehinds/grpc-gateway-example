/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"

	pb "github.com/lukehinds/grpc-auth/gen/go/proto"
	"google.golang.org/grpc"
	"log"

	"github.com/spf13/cobra"
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}

		client := pb.NewAuthClient(conn)

		// Contact the server and print out its response.
		resp, err := client.Login(context.Background(), &pb.LoginRequest{
			Username: "luke",
			Password: "password",
		})

		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}

		fmt.Println(resp)

	},
}

func init() {
	rootCmd.AddCommand(clientCmd)
}
