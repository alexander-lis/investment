package main

import (
	"fmt"
	"net/http"
	"os"
)

var userServerAdress = fmt.Sprintf("%s:%s", os.Getenv("USER_HOST"), os.Getenv("USER_PORT"))

func main() {
	fmt.Print("GATEWAY")
	http.ListenAndServe(":8080", nil)
}

func HTTPProxy(proxyAddr string){

    grpcGwMux:=runtime.NewServeMux()

    // Подключение к сервису User
    grpcUserConn, err:=grpc.Dial(
        userServerAdress,
        grpc.WithInsecure(),
    )
    if err!=nil{
        log.Fatalln("Failed to connect to User service", err)
    }
    defer grpcUserConn.Close()

    err = userService.RegisterUserServiceHandler(
        context.Background(),
        grpcGwMux,
        grpcUserConn,
    )
    if err!=nil{
        log.Fatalln("Failed to start HTTP server", err)
    }


    // Настройка маршрутов с стороны REST
    mux:=http.NewServeMux()
    
    mux.Handle("/api/v1/",grpcGwMux)
    mux.HandleFunc("/",helloworld)

    fmt.Println("starting HTTP server at "+proxyAddr)
    log.Fatal(http.ListenAndServe(proxyAddr,mux))
}
