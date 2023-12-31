package main

import (
	"context"
	"net/http"

	"github.com/alexander-lis/investment/src/app/shared/protobuf/user/authentication"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
)

func main() {
	// Создание контекста и маршрутизатора.
	mux := runtime.NewServeMux()
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Регистрация внутренних gRPC сервисов.
	// Параметры проксирования запросов.
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	// Регистрация AuthenticationService.
	err := authentication.RegisterAuthenticationServiceHandlerFromEndpoint(ctx, mux, "localhost:9000", opts)
	if err != nil {
		grpclog.Fatal(err)
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	http.ListenAndServe(":8081", mux)
}
