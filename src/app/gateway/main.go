package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/alexander-lis/investment/src/app/shared/protobuf/user/authentication"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

var userServerAdress = fmt.Sprintf("%s:%s", os.Getenv("USER_HOST"), os.Getenv("USER_PORT"))

func main() {
	fmt.Print("GATEWAY")
	http.ListenAndServe(":8080", nil)
}

func HTTPProxy(proxyAddr string) {

	grpcGwMux := runtime.NewServeMux()

	// Подключение к сервису User
	grpcUserConn, err := grpc.Dial(
		userServerAdress,
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("Failed to connect to User service", err)
	}
	defer grpcUserConn.Close()

	if err := authentication.RegisterAuthenticationServiceHandler(
		context.Background(),
		grpcGwMux,
		grpcUserConn,
	); err != nil {
		log.Fatalln("Failed to start HTTP server", err)
	}

	// Настройка маршрутов с стороны REST
	mux := http.NewServeMux()

	mux.Handle("/api/", grpcGwMux)

	fmt.Println("starting HTTP server at " + proxyAddr)
	log.Fatal(http.ListenAndServe(proxyAddr, mux))
}
