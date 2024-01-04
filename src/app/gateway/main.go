package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"alexander-lis/investment/shared/infrastructure"
	"alexander-lis/investment/shared/protobuf/user/authentication"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	// Конфигурация конечных точек.
	appPort          = infrastructure.PortFromEnvOrDefault("PORT")
	userServiceUrl   = infrastructure.UrlFromEnvOrDefault("USER_SERVICE_HOST", "USER_SERVICE_PORT")
	stockServiceUrl  = infrastructure.UrlFromEnvOrDefault("STOCK_SERVICE_HOST", "STOCK_SERVICE_PORT")
	budgetServiceUrl = infrastructure.UrlFromEnvOrDefault("BUDGET_SERVICE_HOST", "BUDGET_SERVICE_PORT")
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
		registerGrpcClients(ctx),
	)
}

func logEndpoints() {
	log.Printf("Listening on %s\n", appPort)
	log.Printf("UserService URL: %s", userServiceUrl)
	log.Printf("StockService URL: %s", stockServiceUrl)
	log.Printf("BudgetService URL: %s", budgetServiceUrl)
}

func registerGrpcClients(ctx context.Context) (mux *http.ServeMux) {
	// gwMux = runtime.NewServeMux()
	// gwOpts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	// // TODO: судя по всему проксирование не понадобится, так как нужно делать BFF.
	// if err := authentication.RegisterAuthenticationServiceHandlerFromEndpoint(ctx, gwMux, userServiceUrl, gwOpts); err != nil {
	// 	log.Fatal(err)
	// }

	mux = http.NewServeMux()
	mux.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
		opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
		conn, _ := grpc.Dial(userServiceUrl, opts...)

		client := authentication.NewAuthenticationServiceClient(conn)
		res, _ := client.SignUp(ctx, &authentication.SignUpRequest{})

		fmt.Fprint(w, res.Name)
	})

	return
}

func runGateway(mux *http.ServeMux) {
	mux.HandleFunc("/ready", infrastructure.IsReady)
	mux.HandleFunc("/live", infrastructure.IsAlive)
	if err := http.ListenAndServe(appPort, mux); err != nil {
		log.Fatal(err)
	}
}
