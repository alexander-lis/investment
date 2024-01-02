package infrastructure

import (
	"fmt"
	"net/http"
	"os"
)

func UrlFromEnvOrDefault(hostEnv, portEnv, fallbackUrl string) string {
	url := fmt.Sprintf("%s:%s", os.Getenv(hostEnv), os.Getenv(portEnv))
	if url == "" {
		url = fallbackUrl
	}
	return url
}

func IsReady(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL: %s - OK", r.URL.String())
}

func IsAlive(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL: %s - OK", r.URL.String())
}
