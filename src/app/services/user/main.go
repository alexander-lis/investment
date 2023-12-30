package main

import (
	"fmt"
	"log"
	"net"
	"os"
	handlers "user/api"

	"github.com/alexander-lis/investment/src/app/shared/infrastructure"
	"github.com/alexander-lis/investment/src/app/shared/protobuf/user/authentication"
	"google.golang.org/grpc"
)

func main() {
	ServiceName := "UserService"
	Host := os.Getenv("HOST")
	Port := os.Getenv("PORT")
	LocalPort := os.Getenv("PORT_LOCAL")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(infrastructure.AccessLogInterceptor),
	)

	authentication.AuthenticationServiceServer(grpcServer, &handlers.AuthentiationServerImpl{})

	log.Printf("%s service started on  %s:%s (localhos:%s)", ServiceName, Host, Port, LocalPort)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
