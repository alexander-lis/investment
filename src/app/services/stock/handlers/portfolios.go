package handlers

import (
	"alexander-lis/investment/services/stock/persistence"
	"alexander-lis/investment/shared/infrastructure"
	portfolio "alexander-lis/investment/shared/protobuf/services/stock/proto/v1"
	"context"
	"fmt"
	"time"
)

type PortfolioServiceServerImpl struct {
	portfolio.UnimplementedPortfolioServiceServer
	PortfolioRepository infrastructure.Repository[persistence.Portfolio]
}

func (p *PortfolioServiceServerImpl) CreatePortfolio(ctx context.Context, request *portfolio.CreatePortfolioRequest) (*portfolio.CreatePortfolioResponse, error) {
	from, err := time.Parse(time.RFC3339, request.From)
	if err != nil {
		return nil, fmt.Errorf("Invalid from date")
	}

	to, err := time.Parse(time.RFC3339, request.To)
	if err != nil {
		return nil, fmt.Errorf("Invalid to date")
	}

	p.PortfolioRepository.CreateOrUpdate(&persistence.Portfolio{
		Name: request.Name,
		From: from,
		To:   to,
	})

	return &portfolio.CreatePortfolioResponse{Id: "sds"}, nil
}

func (p *PortfolioServiceServerImpl) GetPortfolios(ctx context.Context, request *portfolio.GetPortfoliosRequest) (*portfolio.GetPortfoliosResponse, error) {
	portfolios := p.PortfolioRepository.ReadAll()

	return &portfolio.GetPortfoliosResponse{
		Portfolios: infrastructure.Map[persistence.Portfolio, portfolio.Portfolio](portfolios, func(p *persistence.Portfolio) *portfolio.Portfolio {
			return &portfolio.Portfolio{
				Name: p.Name,
			}
		}),
	}, nil
}

func (p *PortfolioServiceServerImpl) GetPortfolio(ctx context.Context, request *portfolio.GetPortfolioRequest) (*portfolio.GetPortfolioResponse, error) {
	//TODO implement me
	panic("implement me")
}
