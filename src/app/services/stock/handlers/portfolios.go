package handlers

import (
	"alexander-lis/investment/services/stock/repositories"
	"alexander-lis/investment/shared/infrastructure"
	portfolioProto "alexander-lis/investment/shared/protobuf/services/stock/proto/v1"
	"context"
)

type PortfolioServiceServerImpl struct {
	portfolioProto.UnimplementedPortfolioServiceServer
	PortfolioRepository infrastructure.Repository[repositories.Portfolio]
}

func (p *PortfolioServiceServerImpl) CreatePortfolio(ctx context.Context, request *portfolioProto.CreatePortfolioRequest) (*portfolioProto.CreatePortfolioResponse, error) {
	from, err := infrastructure.TimeFromIsoString(request.From)
	if err != nil {
		return nil, err
	}

	to, err := infrastructure.TimeFromIsoString(request.To)
	if err != nil {
		return nil, err
	}

	id, err := p.PortfolioRepository.Create(&repositories.Portfolio{
		Name: request.Name,
		From: from,
		To:   to,
	})

	if err != nil {
		return nil, err
	}

	return &portfolioProto.CreatePortfolioResponse{
		Portfolio: &portfolioProto.Portfolio{
			Id:   id,
			Name: request.Name,
			From: request.From,
			To:   request.To,
		},
	}, nil
}

func (p *PortfolioServiceServerImpl) GetPortfolios(ctx context.Context, request *portfolioProto.GetPortfoliosRequest) (*portfolioProto.GetPortfoliosResponse, error) {
	result, err := p.PortfolioRepository.ReadAll()
	if err != nil {
		return nil, err
	}

	portfolios := make([]*portfolioProto.Portfolio, 0, len(result))
	for _, p := range result {
		portfolios = append(portfolios,
			&portfolioProto.Portfolio{
				Name: p.Name,
				Id:   p.Id,
				To:   infrastructure.TimeToIsoString(p.To),
				From: infrastructure.TimeToIsoString(p.From),
			},
		)
	}

	return &portfolioProto.GetPortfoliosResponse{Portfolios: portfolios}, nil
}

func (p *PortfolioServiceServerImpl) GetPortfolio(ctx context.Context, request *portfolioProto.GetPortfolioRequest) (*portfolioProto.GetPortfolioResponse, error) {
	result, err := p.PortfolioRepository.Read(request.PortfolioId)
	if err != nil {
		return nil, err
	}

	portfolio := &portfolioProto.Portfolio{Id: result.Id,
		Name: result.Name,
		From: infrastructure.TimeToIsoString(result.From),
		To:   infrastructure.TimeToIsoString(result.To),
	}

	return &portfolioProto.GetPortfolioResponse{Portfolio: portfolio}, nil
}

func (p *PortfolioServiceServerImpl) UpdatePortfolio(ctx context.Context, request *portfolioProto.UpdatePortfolioRequest) (*portfolioProto.Empty, error) {
	from, err := infrastructure.TimeFromIsoString(request.Portfolio.From)
	if err != nil {
		return nil, err
	}

	to, err := infrastructure.TimeFromIsoString(request.Portfolio.To)
	if err != nil {
		return nil, err
	}

	err = p.PortfolioRepository.Update(&repositories.Portfolio{
		Id:   request.Portfolio.Id,
		Name: request.Portfolio.Name,
		From: from,
		To:   to,
	})

	if err != nil {
		return nil, err
	}

	return &portfolioProto.Empty{}, nil
}

func (p *PortfolioServiceServerImpl) DeletePortfolio(ctx context.Context, request *portfolioProto.DeletePortfolioRequest) (*portfolioProto.Empty, error) {
	err := p.PortfolioRepository.Delete(request.PortfolioId)

	if err != nil {
		return nil, err
	}

	return &portfolioProto.Empty{}, nil
}
