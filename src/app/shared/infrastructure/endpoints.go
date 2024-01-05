package infrastructure

import (
	"fmt"
	"net/http"
	"os"
)

func UrlFromEnvOrDefault(hostEnv, portEnv string) string {
	host := os.Getenv(hostEnv)
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv(portEnv)
	if port == "" {
		port = "8080"
	}

	return fmt.Sprintf("%s:%s", host, port)
}

func PortFromEnvOrDefault(portEnv string) string {
	port := os.Getenv(portEnv)
	if port == "" {
		port = "8080"
	}

	return fmt.Sprintf(":%s", port)
}

func IsReady(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL: %s - OK", r.URL.String())
}

func IsAlive(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL: %s - OK", r.URL.String())
}
