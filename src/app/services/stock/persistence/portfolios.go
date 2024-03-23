package persistence

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

const portfoliosCollection string = "portfolios"

// Portfolio entity.
type Portfolio struct {
	Id   string    `bson:"_id,omitempty"`
	Name string    `bson:"name,omitempty"`
	From time.Time `bson:"from,omitempty"`
	To   time.Time `bson:"to,omitempty"`
}

// PortfolioRepository with MongoDB.
type PortfolioRepository struct {
	Context  context.Context
	Database *mongo.Database
}

func (p PortfolioRepository) Create(entity *Portfolio) (string, error) {
	coll := p.Database.Collection(portfoliosCollection)
	result, err := coll.InsertOne(p.Context, entity)
	if err != nil {
		return "", err
	}
	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (p PortfolioRepository) Update(entity *Portfolio) error {
	fmt.Print(entity)
	//coll := p.Database.Collection(portfoliosCollection)
	//result, err := coll.UpdateOne(p.Context, entity)
	//if err != nil {
	//	return "", err
	//}
	//return result.InsertedID.(primitive.ObjectID).Hex(), nil
	return nil
}

func (p PortfolioRepository) Read(id string) (Portfolio, error) {
	//TODO implement me
	panic("implement me")
}

func (p PortfolioRepository) ReadAll() ([]Portfolio, error) {
	coll := p.Database.Collection(portfoliosCollection)

	cursor, err := coll.Find(p.Context, bson.D{})
	if err != nil {
		return nil, err
	}

	var results []Portfolio
	if err = cursor.All(p.Context, &results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p PortfolioRepository) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}
