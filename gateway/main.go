package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Print("GATEWAY")
	http.ListenAndServe(":8080", nil)
}
