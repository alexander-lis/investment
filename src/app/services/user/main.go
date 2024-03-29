package main

import (
	"log"
	"net"

	"alexander-lis/investment/shared/infrastructure"
	user "alexander-lis/investment/shared/protobuf/services/user/proto/v1"

	handers "alexander-lis/investment/services/user/handlers"

	"google.golang.org/grpc"
)

var (
	// Конфигурация конечных точек.
	appPort = infrastructure.PortFromEnvOrDefault("PORT", "9092")
)

func main() {
	// Лог параметров.
	logEndpoints()

	// Конфигурация сервера
	grpcServer := configureGrpcServer()

	// Регистрация обработчиков.
	registerHandlers(grpcServer)

	// Запуск сервера.
	startGrpcServer(grpcServer)
}

func logEndpoints() {
	log.Printf("Listening on %s\n", appPort)
}

func configureGrpcServer() (grpcServer *grpc.Server) {
	grpcServer = grpc.NewServer(
	// 	grpc.UnaryInterceptor(infrastructure.AccessLogInterceptor)
	)

	return
}

func registerHandlers(grpcServer *grpc.Server) {
	user.RegisterAuthenticationServiceServer(grpcServer, &handers.AuthenticationServiceServerImpl{})
}

func startGrpcServer(grpcServer *grpc.Server) {
	listener, err := net.Listen("tcp", appPort)
	if err != nil {
		log.Fatal(err)
	}

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
