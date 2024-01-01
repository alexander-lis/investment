package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"alexander-lis/investment/shared/infrastructure"
	"alexander-lis/investment/shared/protobuf/user/authentication"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Загрузка конфигурации.
	infrastructure.BuildConfiguration()

	// Создание контекста.
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Параметры проксирования запросов.
	gwMux := runtime.NewServeMux()
	gwOpts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	// Регистрация AuthenticationService.
	if err := authentication.RegisterAuthenticationServiceHandlerFromEndpoint(ctx, gwMux, infrastructure.Hosts.UserServiceUrl, gwOpts); err != nil {
		log.Fatal(err)
	}

	log.Printf("Running on :%s", infrastructure.Hosts.GatewayPort)

	// Запуск сервера и передача проксирующеого мультиплексора.
	mux := http.NewServeMux()
	mux.Handle("/api/", gwMux)
	mux.HandleFunc("/ready", isReady)
	mux.HandleFunc("/live", isAlive)
	if err := http.ListenAndServe(infrastructure.ServiceUrl("", "80"), mux); err != nil {
		log.Fatal(err)
	}
}

func isReady(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL: %s - OK", r.URL.String())
}

func isAlive(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL: %s - OK", r.URL.String())
}
