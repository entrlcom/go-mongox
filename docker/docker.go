package docker

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	client    *mongo.Client
	container *mongodb.MongoDBContainer
}

func NewMongo(ctx context.Context, opts ...MongoOption) (Mongo, error) {
	ops := make([]testcontainers.ContainerCustomizer, 0, len(opts)+1)

	ops = append(ops, WithMongoVersion(""))

	for _, opt := range opts {
		ops = append(ops, opt)
	}

	container, err := mongodb.RunContainer(ctx, ops...)
	if err != nil {
		return Mongo{}, err
	}

	uri, err := container.ConnectionString(ctx)
	if err != nil {
		return Mongo{}, err
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return Mongo{}, err
	}

	mongoDocker := Mongo{
		client:    client,
		container: container,
	}

	return mongoDocker, nil
}

func (x Mongo) GetClient() *mongo.Client {
	return x.client
}

func (x Mongo) Shutdown(ctx context.Context) error {
	return x.container.Terminate(ctx)
}

type MongoOption testcontainers.ContainerCustomizer

func WithMongoPassword(password string) MongoOption {
	return mongodb.WithPassword(password)
}

func WithMongoUsername(username string) MongoOption {
	return mongodb.WithUsername(username)
}

func WithMongoVersion(version string) MongoOption {
	if len(version) == 0 {
		version = "latest"
	}

	return testcontainers.WithImage("mongo:" + version)
}
