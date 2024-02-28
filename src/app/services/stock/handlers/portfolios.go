package handlers

import (
	"alexander-lis/investment/services/stock/persistence"
	"alexander-lis/investment/shared/infrastructure"
	portfolio "alexander-lis/investment/shared/protobuf/services/stock/proto/v1"
	"context"
)

type PortfolioServiceServerImpl struct {
	portfolio.UnimplementedPortfolioServiceServer
	PortfolioRepository infrastructure.Repository[persistence.Portfolio]
}

func (p *PortfolioServiceServerImpl) GetPortfolios(ctx context.Context, request *portfolio.GetPortfoliosRequest) (*portfolio.GetPortfoliosResponse, error) {
	portfolios := p.PortfolioRepository.ReadAll()

	dtos := infrastructure.Map[persistence.Portfolio, portfolio.Portfolio](portfolios, persistence.PortfoioToDto)

	return &portfolio.GetPortfoliosResponse{
		Portfolios: dtos,
	}, nil
}

func (p *PortfolioServiceServerImpl) GetPortfolio(ctx context.Context, request *portfolio.GetPortfolioRequest) (*portfolio.GetPortfolioResponse, error) {
	//TODO implement me
	panic("implement me")
}
