package main

import (
	"log"
	"net"

	"alexander-lis/investment/services/user/api"
	"alexander-lis/investment/shared/infrastructure"
	"alexander-lis/investment/shared/protobuf/user/authentication"

	"google.golang.org/grpc"
)

func main() {
	// Загрузка конфигурации.
	infrastructure.BuildConfiguration()

	// gRPC - конфигурация сервера.
	grpcServer := grpc.NewServer(
	//grpc.UnaryInterceptor(infrastructure.AccessLogInterceptor),
	)

	// gRPC - регистрация сервисов.
	authentication.RegisterAuthenticationServiceServer(grpcServer, &api.AuthentiationServiceServerImpl{})

	// gRPC - запуск сервера.
	listener, err := net.Listen("tcp", infrastructure.ServiceUrl("", infrastructure.Hosts.UserServicePort))
	if err != nil {
		log.Fatal(err)
	}

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}

}
