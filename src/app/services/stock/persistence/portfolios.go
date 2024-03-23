package persistence

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

const portfoliosCollection string = "portfolios"

// Portfolio entity.
type Portfolio struct {
	Id   primitive.ObjectID `bson:"_id"`
	Name string             `bson:"name,omitempty"`
	From time.Time          `bson:"from,omitempty"`
	To   time.Time          `bson:"to,omitempty"`
}

// PortfolioRepository with MongoDB.
type PortfolioRepository struct {
	Context  context.Context
	Database *mongo.Database
}

func (p PortfolioRepository) CreateOrUpdate(entity *Portfolio) {
	coll := p.Database.Collection(portfoliosCollection)
	result, err := coll.InsertOne(p.Context, entity)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(result)
}

func (p PortfolioRepository) Read(id string) *Portfolio {
	//TODO implement me
	panic("implement me")
}

func (p PortfolioRepository) ReadAll() []*Portfolio {
	//coll := p.Database.Collection(portfoliosCollection)
	//
	//filter := bson.D{}
	//
	//cursor, err := coll.Find(p.Context, filter)
	//if err != nil {
	//	panic(err)
	//}
	//
	//var results []Portfolio
	//if err = cursor.All(p.Context, &results); err != nil {
	//	panic(err)
	//}
	//
	//// Prints the results of the find operation as structs
	//for _, result := range results {
	//	cursor.Decode(&result)
	//	output, err := json.MarshalIndent(result, "", "    ")
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Printf("%s\n", output)
	//}
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
