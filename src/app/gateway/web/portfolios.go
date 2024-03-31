package web

import (
	"alexander-lis/investment/shared/infrastructure"
	stock "alexander-lis/investment/shared/protobuf/services/stock/proto/v1"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

// Data Transfer Objects.
type PortfolioDto struct {
	Name string `json:"name"`
}

// RegisterPortfolios registers portfolios endpoints.
func RegisterPortfolios(grpcClients *infrastructure.GrpcServiceClients, e *echo.Echo) {
	e.POST("/portfolios", createPortfolio(grpcClients))
	e.GET("/portfolios", getPortfolios(grpcClients))
	e.GET("/portfolios/:id", getPortfolio(grpcClients))
	e.PUT("/portfolios/:id", updatePortfolio(grpcClients))
	e.DELETE("/portfolios/:id", deletePortfolio(grpcClients))
}

func createPortfolio(grpcClients *infrastructure.GrpcServiceClients) func(c echo.Context) error {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "")
	}
}

func getPortfolios(grpcClients *infrastructure.GrpcServiceClients) func(c echo.Context) error {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		response, error := grpcClients.PortfolioServiceClient.GetPortfolios(ctx, &stock.GetPortfoliosRequest{})
		if error != nil {
			return c.String(http.StatusInternalServerError, error.Error())
		}

		return c.String(http.StatusOK, fmt.Sprintf("%s", response))
	}
}

func getPortfolio(grpcClients *infrastructure.GrpcServiceClients) func(c echo.Context) error {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "")
	}
}

func updatePortfolio(grpcClients *infrastructure.GrpcServiceClients) func(c echo.Context) error {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "")
	}
}

func deletePortfolio(grpcClients *infrastructure.GrpcServiceClients) func(c echo.Context) error {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "")
	}
}
