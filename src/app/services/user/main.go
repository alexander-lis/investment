package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"user/api"

	"github.com/alexander-lis/investment/src/app/shared/protobuf/user/authentication"
	"google.golang.org/grpc"
)

func main() {
	ServiceName := "UserService"
	Host := os.Getenv("HOST")
	if Host == "" {
		Host = "localhost"
	}
	Port := os.Getenv("PORT")
	if Port == "" {
		Port = "9000"
	}
	LocalPort := os.Getenv("PORT_LOCAL")
	if LocalPort == "" {
		LocalPort = "9000"
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(
	//grpc.UnaryInterceptor(infrastructure.AccessLogInterceptor),
	)

	authentication.RegisterAuthenticationServiceServer(grpcServer, &api.AuthentiationServiceServerImpl{})

	log.Printf("%s service started on  %s:%s (localhost:%s)", ServiceName, Host, Port, LocalPort)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
