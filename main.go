package main

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"entrlcom.dev/mongox/collection"
)

type Person struct {
	DataOfBirth time.Time `bson:"data_of_birth,omitempty"`
	Name        string    `bson:"name,omitempty"`
	TimeCreated time.Time `bson:"time_created,omitempty"`
}

type Repository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func (x Repository) Create(ctx context.Context, person Person) error {
	return collection.Insert[Person](x.collection).Insert(ctx, person)
}

func NewRepository(client *mongo.Client) Repository {
	return Repository{
		client:     client,
		collection: client.Database("example").Collection("person"),
	}
}

func main() {
	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		// TODO: Handle error.
		return
	}

	repository := NewRepository(client)

	if err := repository.Create(ctx, Person{
		DataOfBirth: time.Date(2020, time.May, 15, 0, 0, 0, 0, time.UTC),
		Name:        "John Doe",
		TimeCreated: time.Now().UTC(),
	}); err != nil {
		// TODO: Handle error.
		return
	}

	// ...
}
