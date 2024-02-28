package infrastructure

type Entity interface {
}

type MongoRepository struct {
	MongoUrl string
}

type Repository[TEntity Entity] interface {
	CreateOrUpdate(entity *TEntity)
	Read(id string) *TEntity
	ReadAll() []*TEntity
	Delete(id string)
}
