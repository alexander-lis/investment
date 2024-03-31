package main

import (
	"alexander-lis/investment/gateway/web"
	"alexander-lis/investment/shared/infrastructure"
	stock "alexander-lis/investment/shared/protobuf/services/stock/proto/v1"
	user "alexander-lis/investment/shared/protobuf/services/user/proto/v1"
	"context"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
	"log"
)

var (
	config = &infrastructure.Config{
		AppPort:          infrastructure.PortFromEnvOrDefault("PORT", "9090"),
		UserServiceUrl:   infrastructure.UrlFromEnvOrDefault("USER_SERVICE_HOST", "USER_SERVICE_PORT", "localhost", "9092"),
		StockServiceUrl:  infrastructure.UrlFromEnvOrDefault("STOCK_SERVICE_HOST", "STOCK_SERVICE_PORT", "localhost", "9093"),
		BudgetServiceUrl: infrastructure.UrlFromEnvOrDefault("BUDGET_SERVICE_HOST", "BUDGET_SERVICE_PORT", "localhost", "9094"),
	}
)

func main() {
	// Лог параметров.
	logEndpoints()

	// Контекст.
	ctx := context.Background()
	ctx, cancelCtx := context.WithCancel(ctx)
	defer cancelCtx()

	// Регистрация gRPC клиентов.
	grpcClients := registerGrpcClients(ctx)

	// Запуск сервера.
	runServer(grpcClients)
}

func logEndpoints() {
	log.Printf("Listening on %s\n", config.AppPort)
	log.Printf("UserService URL: %s", config.UserServiceUrl)
	log.Printf("StockService URL: %s", config.StockServiceUrl)
	log.Printf("BudgetService URL: %s", config.BudgetServiceUrl)
}

func registerGrpcClients(ctx context.Context) *infrastructure.GrpcServiceClients {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	stockConn, err := grpc.DialContext(ctx, config.StockServiceUrl, opts...)
	if err != nil {
		log.Fatal(err)
	}
	userConn, err := grpc.DialContext(ctx, config.UserServiceUrl, opts...)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		<-ctx.Done()
		if cerr := stockConn.Close(); cerr != nil {
			grpclog.Infof("Failed to close conn to %s: %v", config.StockServiceUrl, cerr)
		}
		if cerr := userConn.Close(); cerr != nil {
			grpclog.Infof("Failed to close conn to %s: %v", config.UserServiceUrl, cerr)
		}
	}()

	return &infrastructure.GrpcServiceClients{
		PortfolioServiceClient:      stock.NewPortfolioServiceClient(stockConn),
		AccountServiceClient:        user.NewAccountServiceClient(userConn),
		AuthenticationServiceClient: user.NewAuthenticationServiceClient(userConn),
	}
}

func runServer(grpcClients *infrastructure.GrpcServiceClients) {
	e := echo.New()
	web.RegisterPortfolios(grpcClients, e)
	e.Logger.Fatal(e.Start(config.AppPort))
}
