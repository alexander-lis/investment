package persistence

import (
	"alexander-lis/investment/shared/infrastructure"
	portfolio "alexander-lis/investment/shared/protobuf/services/stock/proto/v1"
	"time"
)

// Portfolio entity.
type Portfolio struct {
	Name     string
	From, To time.Time
}

// PortfoioToDto maps entity to dto.
func PortfoioToDto(p *Portfolio) *portfolio.Portfolio {
	return &portfolio.Portfolio{
		Name: p.Name,
	}
}

// PortfolioRepository for MongoDB.
type PortfolioRepository struct {
	infrastructure.MongoRepository
}

func (p PortfolioRepository) CreateOrUpdate(entity *Portfolio) {
	//TODO implement me
	panic("implement me")
}

func (p PortfolioRepository) Read(id string) *Portfolio {
	//TODO implement me
	panic("implement me")
}

func (p PortfolioRepository) ReadAll() []*Portfolio {
	portfolios := []*Portfolio{
		{
			Name: "First",
			From: time.Time{},
			To:   time.Time{},
		},
		{
			Name: "Second",
			From: time.Time{},
			To:   time.Time{},
		},
	}
	return portfolios
}

func (p PortfolioRepository) Delete(id string) {
	//TODO implement me
	panic("implement me")
}
