package collection

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
)

type FindManyOperation[T any] struct { //nolint:govet // OK.
	collection *mongo.Collection
	filter     bsoncore.Document
	options    *options.FindOptions
}

func (x FindManyOperation[T]) FindMany(ctx context.Context) ([]T, error) {
	cursor, err := x.collection.Find(ctx, x.filter, x.options)
	if err != nil {
		return nil, err
	}

	defer func() { _ = cursor.Close(ctx) }() //nolint:errcheck // OK.

	var limit int64

	if x.options != nil && x.options.Limit != nil {
		limit = *x.options.Limit
	}

	documents := make([]T, 0, limit)

	for cursor.Next(ctx) {
		var document T

		if err := cursor.Decode(&document); err != nil {
			return nil, err
		}

		documents = append(documents, document)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(documents) == 0 {
		return nil, mongo.ErrNoDocuments
	}

	return documents, nil
}

func FindMany[T any](c *mongo.Collection, opts ...FindManyOption[T]) FindManyOperation[T] {
	x := FindManyOperation[T]{
		collection: c,
		filter:     bsoncore.NewDocumentBuilder().Build(),
		options:    nil,
	}

	for _, opt := range opts {
		opt(&x)
	}

	return x
}

type FindManyOption[T any] func(x *FindManyOperation[T])

func WithFindManyFilter[T any](filter bsoncore.Document) FindManyOption[T] {
	return func(x *FindManyOperation[T]) {
		x.filter = filter
	}
}

func WithFindManyOptions[T any](opt *options.FindOptions) FindManyOption[T] {
	return func(x *FindManyOperation[T]) {
		x.options = opt
	}
}
