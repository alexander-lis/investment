package infrastructure

import (
	"fmt"
	"os"
)

var Hosts struct {
	GatewayPort,
	UserServiceHost,
	UserServicePort,
	UserServiceUrl,
	PlanningServiceHost,
	PlanningServicePort,
	PlanningServiceUrl string
}

func BuildConfiguration() {
	// Gateway.
	if Hosts.GatewayPort = os.Getenv("GATEWAY_PORT"); Hosts.GatewayPort == "" {
		Hosts.GatewayPort = "80"
	}

	// User service.
	if Hosts.UserServiceHost = os.Getenv("USER_SERVICE_HOST"); Hosts.UserServiceHost == "" {
		Hosts.UserServiceHost = "localhost"
	}
	if Hosts.UserServicePort = os.Getenv("USER_SERVICE_PORT"); Hosts.UserServicePort == "" {
		Hosts.UserServicePort = "9000"
	}
	Hosts.UserServiceUrl = ServiceUrl(Hosts.UserServiceHost, Hosts.UserServicePort)

	// Planning service.
	if Hosts.PlanningServiceHost = os.Getenv("PLANNING_SERVICE_HOST"); Hosts.PlanningServiceHost == "" {
		Hosts.PlanningServiceHost = "localhost"
	}
	if Hosts.PlanningServicePort = os.Getenv("PLANNING_SERVICE_HOST"); Hosts.PlanningServicePort == "" {
		Hosts.PlanningServicePort = "9001"
	}
	Hosts.PlanningServiceUrl = ServiceUrl(Hosts.PlanningServiceHost, Hosts.PlanningServicePort)
}

func ServiceUrl(host, port string) string {
	return fmt.Sprintf("%s:%s", host, port)
}
