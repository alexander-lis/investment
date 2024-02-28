package main

import (
	"alexander-lis/investment/services/stock/handlers"
	"alexander-lis/investment/services/stock/persistence"
	portfolio "alexander-lis/investment/shared/protobuf/services/stock/proto/v1"
	"log"
	"net"

	"alexander-lis/investment/shared/infrastructure"

	"google.golang.org/grpc"
)

var (
	// Конфигурация конечных точек.
	appPort  = infrastructure.PortFromEnvOrDefault("PORT", "9093")
	mongoUrl = infrastructure.MongoUrlFromEnv()
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
	portfolio.RegisterPortfolioServiceServer(grpcServer, &handlers.PortfolioServiceServerImpl{
		PortfolioRepository: persistence.PortfolioRepository{
			MongoRepository: infrastructure.MongoRepository{MongoUrl: mongoUrl},
		},
	})
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
