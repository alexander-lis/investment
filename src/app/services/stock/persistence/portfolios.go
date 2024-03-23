package persistence

import (
	portfolio "alexander-lis/investment/shared/protobuf/services/stock/proto/v1"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

const collection string = "portfolios"

// Portfolio entity.
type Portfolio struct {
	Id   string    `bson:"id,omitempty"`
	Name string    `bson:"name,omitempty"`
	From time.Time `bson:"from,omitempty"`
	To   time.Time `bson:"to,omitempty"`
}

// PortfolioToDto maps entity to dto.
func PortfolioToDto(p *Portfolio) *portfolio.Portfolio {
	return &portfolio.Portfolio{
		Name: p.Name,
	}
}

// PortfolioRepository for MongoDB.
type PortfolioRepository struct {
	Client *mongo.Client
}

func (p PortfolioRepository) CreateOrUpdate(entity *Portfolio) {

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
