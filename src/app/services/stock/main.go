package main

import (
	"alexander-lis/investment/services/stock/handlers"
	"alexander-lis/investment/services/stock/persistence"
	"alexander-lis/investment/shared/infrastructure"
	portfolio "alexander-lis/investment/shared/protobuf/services/stock/proto/v1"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	// Конфигурация конечных точек.
	appPort  = infrastructure.PortFromEnvOrDefault("PORT", "9093")
	mongoUri = infrastructure.MongoUriFromEnv()
)

func main() {
	// Контекст и обработка завершения.
	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
	}()

	// Лог параметров.
	logEndpoints()

	// Конфигурация сервера
	grpcServer := configureGrpcServer()

	// Регистрация обработчиков.
	registerHandlers(grpcServer, ctx)

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

func registerHandlers(grpcServer *grpc.Server, ctx context.Context) {
	mongoClient := infrastructure.GetMongoClient(mongoUri)

	portfolio.RegisterPortfolioServiceServer(grpcServer, &handlers.PortfolioServiceServerImpl{
		PortfolioRepository: persistence.PortfolioRepository{
			Database: mongoClient.Database("stock"),
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
