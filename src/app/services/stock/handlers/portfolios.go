package handlers

import (
	"alexander-lis/investment/services/stock/persistence"
	"alexander-lis/investment/shared/infrastructure"
	portfolio "alexander-lis/investment/shared/protobuf/services/stock/proto/v1"
	"context"
	"time"
)

type PortfolioServiceServerImpl struct {
	portfolio.UnimplementedPortfolioServiceServer
	PortfolioRepository infrastructure.Repository[persistence.Portfolio]
}

func (p *PortfolioServiceServerImpl) CreatePortfolio(ctx context.Context, request *portfolio.CreatePortfolioRequest) (*portfolio.CreatePortfolioResponse, error) {
	from, err := time.Parse(time.RFC3339, request.From)
	if err != nil {
		return nil, err
	}

	to, err := time.Parse(time.RFC3339, request.To)
	if err != nil {
		return nil, err
	}

	id, err := p.PortfolioRepository.Create(&persistence.Portfolio{
		Name: request.Name,
		From: from,
		To:   to,
	})

	if err != nil {
		return nil, err
	}

	return &portfolio.CreatePortfolioResponse{
		Portfolio: &portfolio.Portfolio{
			Id:   id,
			Name: request.Name,
			From: request.From,
			To:   request.To,
		},
	}, nil
}

func (p *PortfolioServiceServerImpl) GetPortfolios(ctx context.Context, request *portfolio.GetPortfoliosRequest) (*portfolio.GetPortfoliosResponse, error) {
	result, err := p.PortfolioRepository.ReadAll()
	if err != nil {
		return nil, err
	}

	portfolios := make([]*portfolio.Portfolio, 0, len(result))
	for _, p := range result {
		portfolios = append(portfolios,
			&portfolio.Portfolio{
				Name: p.Name,
				Id:   p.Id,
				To:   p.To.Format(time.RFC3339),
				From: p.From.Format(time.RFC3339),
			},
		)
	}

	return &portfolio.GetPortfoliosResponse{Portfolios: portfolios}, nil
}

func (p *PortfolioServiceServerImpl) GetPortfolio(ctx context.Context, request *portfolio.GetPortfolioRequest) (*portfolio.GetPortfolioResponse, error) {
	//TODO implement me
	panic("implement me")
}
