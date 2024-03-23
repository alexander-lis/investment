package infrastructure

import (
	"fmt"
	"net/http"
	"os"
)

func UrlFromEnvOrDefault(hostEnv, portEnv, defaultHost, defaultPort string) string {
	host := os.Getenv(hostEnv)
	if host == "" {
		host = defaultHost
	}

	port := os.Getenv(portEnv)
	if port == "" {
		port = defaultPort
	}

	return fmt.Sprintf("%s:%s", host, port)
}

func PortFromEnvOrDefault(portEnv, defaultPort string) string {
	port := os.Getenv(portEnv)
	if port == "" {
		port = defaultPort
	}

	return fmt.Sprintf(":%s", port)
}

func IsReady(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL: %s - OK", r.URL.String())
}

func IsAlive(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL: %s - OK", r.URL.String())
}

func MongoUriFromEnv() string {
	return os.Getenv("MONGO_URI")
}
