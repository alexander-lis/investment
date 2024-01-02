package main

import (
	"context"
	"log"
	"net/http"

	"alexander-lis/investment/shared/infrastructure"
	"alexander-lis/investment/shared/protobuf/user/authentication"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	// Конфигурация конечных точек.
	appPort          = infrastructure.UrlFromEnvOrDefault("", "PORT", "")
	userServiceUrl   = infrastructure.UrlFromEnvOrDefault("USER_SERVICE_HOST", "USER_SERVICE_PORT", "localhost:9001")
	stockServiceUrl  = infrastructure.UrlFromEnvOrDefault("STOCK_SERVICE_HOST", "STOCK_SERVICE_PORT", "localhost:9002")
	budgetServiceUrl = infrastructure.UrlFromEnvOrDefault("BUDGET_SERVICE_HOST", "BUDGET_SERVICE_PORT", "localhost:9003")
)

func main() {
	// Лог параметров.
	logEndpoints()

	// Создание контекста.
	ctx := context.Background()
	ctx, cancelCtx := context.WithCancel(ctx)
	defer cancelCtx()

	// Регистрация сервисов и запуск шлюза.
	runGateway(
		registerServices(ctx),
	)
}

func logEndpoints() {
	log.Printf("Listening on %s\n", appPort)
	log.Printf("UserService URL: %s", userServiceUrl)
	log.Printf("StockService URL: %s", stockServiceUrl)
	log.Printf("BudgetService URL: %s", budgetServiceUrl)
}

func registerServices(ctx context.Context) (gwMux *runtime.ServeMux) {
	gwMux = runtime.NewServeMux()
	gwOpts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	if err := authentication.RegisterAuthenticationServiceHandlerFromEndpoint(ctx, gwMux, userServiceUrl, gwOpts); err != nil {
		log.Fatal(err)
	}

	return
}

func runGateway(gwMux *runtime.ServeMux) {
	mux := http.NewServeMux()
	mux.Handle("/api/", gwMux)
	mux.HandleFunc("/ready", infrastructure.IsReady)
	mux.HandleFunc("/live", infrastructure.IsAlive)
	if err := http.ListenAndServe(appPort, mux); err != nil {
		log.Fatal(err)
	}
}
