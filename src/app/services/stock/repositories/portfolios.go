package repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

const portfoliosCollection string = "portfolios"

type Portfolio struct {
	Id   string    `bson:"_id,omitempty"`
	Name string    `bson:"name,omitempty"`
	From time.Time `bson:"from,omitempty"`
	To   time.Time `bson:"to,omitempty"`
}

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
	coll := p.Database.Collection(portfoliosCollection)

	id, err := primitive.ObjectIDFromHex(entity.Id)
	if err != nil {
		return err
	}

	update := bson.D{
		{"$set", bson.D{
			{"name", entity.Name},
			{"from", entity.To},
			{"to", entity.From}},
		},
	}

	_, err = coll.UpdateByID(p.Context, id, update)
	if err != nil {
		return err
	}
	return nil
}

func (p PortfolioRepository) Read(idStr string) (*Portfolio, error) {
	coll := p.Database.Collection(portfoliosCollection)

	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return nil, err
	}

	filter := bson.D{{"_id", id}}

	var portfolio Portfolio
	err = coll.FindOne(p.Context, filter).Decode(&portfolio)
	if err != nil {
		return nil, err
	}

	return &portfolio, nil
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

func (p PortfolioRepository) Delete(idStr string) error {
	coll := p.Database.Collection(portfoliosCollection)

	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return err
	}

	filter := bson.D{{"_id", id}}

	_, err = coll.DeleteOne(p.Context, filter)
	if err != nil {
		return err
	}

	return nil
}
