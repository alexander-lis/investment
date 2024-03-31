package infrastructure

import (
	stock "alexander-lis/investment/shared/protobuf/services/stock/proto/v1"
	user "alexander-lis/investment/shared/protobuf/services/user/proto/v1"
)

type Config struct {
	AppPort          string
	UserServiceUrl   string
	StockServiceUrl  string
	BudgetServiceUrl string
}

type GrpcServiceClients struct {
	// Stock microservice.
	PortfolioServiceClient stock.PortfolioServiceClient

	// User microservice.
	AccountServiceClient        user.AccountServiceClient
	AuthenticationServiceClient user.AuthenticationServiceClient
}
