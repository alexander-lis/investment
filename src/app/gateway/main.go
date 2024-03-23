package main

import (
	"context"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"net/http"

	"alexander-lis/investment/shared/infrastructure"
	stock "alexander-lis/investment/shared/protobuf/services/stock/proto/v1"
	user "alexander-lis/investment/shared/protobuf/services/user/proto/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	// Конфигурация конечных точек.
	appPort          = infrastructure.PortFromEnvOrDefault("PORT", "9090")
	userServiceUrl   = infrastructure.UrlFromEnvOrDefault("USER_SERVICE_HOST", "USER_SERVICE_PORT", "localhost", "9092")
	stockServiceUrl  = infrastructure.UrlFromEnvOrDefault("STOCK_SERVICE_HOST", "STOCK_SERVICE_PORT", "localhost", "9093")
	budgetServiceUrl = infrastructure.UrlFromEnvOrDefault("BUDGET_SERVICE_HOST", "BUDGET_SERVICE_PORT", "localhost", "9094")
)

func main() {
	// Лог параметров.
	logEndpoints()

	// Создание контекста.
	ctx := context.Background()
	ctx, cancelCtx := context.WithCancel(ctx)
	defer cancelCtx()

	// Регистрация сервисов и запуск шлюза.
	gwMux := getGatewayMux(ctx)
	runGateway(gwMux)
}

func logEndpoints() {
	log.Printf("Listening on %s\n", appPort)
	log.Printf("UserService URL: %s", userServiceUrl)
	log.Printf("StockService URL: %s", stockServiceUrl)
	log.Printf("BudgetService URL: %s", budgetServiceUrl)
}

func getGatewayMux(ctx context.Context) (gwMux *runtime.ServeMux) {
	gwMux = runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseProtoNames: true,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true,
			},
		}))
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	stock.RegisterPortfolioServiceHandlerFromEndpoint(ctx, gwMux, stockServiceUrl, opts)
	user.RegisterAuthenticationServiceHandlerFromEndpoint(ctx, gwMux, userServiceUrl, opts)

	return
}

func runGateway(gwMux *runtime.ServeMux) {
	mux := http.NewServeMux()
	mux.HandleFunc("/ready", infrastructure.IsReady)
	mux.HandleFunc("/live", infrastructure.IsAlive)
	mux.Handle("/api/", gwMux)
	if err := http.ListenAndServe(appPort, mux); err != nil {
		log.Fatal(err)
	}
}
