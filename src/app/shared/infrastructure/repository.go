package infrastructure

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

/* Base structures */

type Entity interface {
}

type Repository[TEntity Entity] interface {
	Create(entity *TEntity) (string, error)
	Update(entity *TEntity) error
	Read(id string) (*TEntity, error)
	ReadAll() ([]TEntity, error)
	Delete(id string) error
}

// GetMongoClient returns new MongoDB client.
func GetMongoClient(mongoUri string) *mongo.Client {
	ctx := context.TODO()

	opts := options.Client()
	if mongoUri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable")
	}
	client, err := mongo.Connect(ctx, opts.ApplyURI(mongoUri))
	if err != nil {
		panic(err)
	}

	//if errClose := client.Disconnect(ctx); errClose != nil {
	//	panic(errClose)
	//} else {
	//	log.Printf("Mongo client disconnected")
	//}

	return client
}
