package endpoints

import (
	"alexander-lis/investment/gateway/web/dto"
	"alexander-lis/investment/shared/infrastructure"
	stock "alexander-lis/investment/shared/protobuf/services/stock/proto/v1"
	"github.com/labstack/echo/v4"
	"net/http"
)

// RegisterPortfolios registers portfolios endpoints.
func RegisterPortfolios(grpcClients *infrastructure.GrpcServiceClients, e *echo.Echo) {
	e.POST("/portfolios", createPortfolio(grpcClients))
	e.PATCH("/portfolios/:id", updatePortfolio(grpcClients))
	e.GET("/portfolios", getPortfolios(grpcClients))
	e.GET("/portfolios/:id", getPortfolioFull(grpcClients))
	e.DELETE("/portfolios/:id", deletePortfolio(grpcClients))
}

func createPortfolio(grpcClients *infrastructure.GrpcServiceClients) func(c echo.Context) error {
	return func(c echo.Context) error {
		// Request.
		var createPortfolio dto.CreatePortfolio
		if err := c.Bind(createPortfolio); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		// Service.
		ctx := c.Request().Context()
		response, err := grpcClients.PortfolioServiceClient.CreatePortfolio(ctx, &stock.CreatePortfolioRequest{
			Name: createPortfolio.Name,
			From: createPortfolio.From,
			To:   createPortfolio.To,
		})

		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		// Response.
		return c.JSON(http.StatusOK, response)
	}
}

func updatePortfolio(grpcClients *infrastructure.GrpcServiceClients) func(c echo.Context) error {
	return func(c echo.Context) error {
		// Request.
		id := c.Param("id")

		updatePortfolio := &dto.UpdatePortfolio{}
		if err := c.Bind(updatePortfolio); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		// Service.
		ctx := c.Request().Context()
		_, error := grpcClients.PortfolioServiceClient.UpdatePortfolio(ctx, &stock.UpdatePortfolioRequest{
			Portfolio: &stock.Portfolio{
				Id:   id,
				Name: updatePortfolio.Name,
				To:   updatePortfolio.To,
				From: updatePortfolio.From,
			},
		})
		if error != nil {
			return c.String(http.StatusInternalServerError, error.Error())
		}

		// Response.
		return c.NoContent(http.StatusOK)
	}
}

func getPortfolios(grpcClients *infrastructure.GrpcServiceClients) func(c echo.Context) error {
	return func(c echo.Context) error {
		// Service.
		ctx := c.Request().Context()
		response, error := grpcClients.PortfolioServiceClient.GetPortfolios(ctx, &stock.GetPortfoliosRequest{})
		if error != nil {
			return c.String(http.StatusInternalServerError, error.Error())
		}

		var portfolios []dto.PortfolioPartial
		for _, p := range response.Portfolios {
			portfolios = append(portfolios, dto.PortfolioPartial{
				Id:   p.Id,
				Name: p.Name,
				To:   p.To,
				From: p.From,
			})
		}

		// Response.
		return c.JSON(http.StatusOK, portfolios)
	}
}

func getPortfolioFull(grpcClients *infrastructure.GrpcServiceClients) func(c echo.Context) error {
	return func(c echo.Context) error {
		// Request.
		id := c.Param("id")

		// Service.
		ctx := c.Request().Context()
		response, error := grpcClients.PortfolioServiceClient.GetPortfolio(ctx, &stock.GetPortfolioRequest{PortfolioId: id})
		if error != nil {
			return c.String(http.StatusInternalServerError, error.Error())
		}

		// Response.
		return c.JSON(http.StatusOK, &dto.PortfolioFull{
			Id:      response.Portfolio.Id,
			Name:    response.Portfolio.Name,
			From:    response.Portfolio.From,
			To:      response.Portfolio.To,
			Asset:   make([]dto.Asset, 0),
			Sectors: make([]dto.Sector, 0),
			Groups:  make([]dto.Group, 0),
		})
	}
}

func deletePortfolio(grpcClients *infrastructure.GrpcServiceClients) func(c echo.Context) error {
	return func(c echo.Context) error {
		// Request.
		id := c.Param("id")

		// Service.
		ctx := c.Request().Context()
		_, error := grpcClients.PortfolioServiceClient.DeletePortfolio(ctx, &stock.DeletePortfolioRequest{PortfolioId: id})
		if error != nil {
			return c.String(http.StatusInternalServerError, error.Error())
		}

		// Response.
		return c.NoContent(http.StatusOK)
	}
}
