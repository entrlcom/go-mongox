package collection

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
)

type FindOperation[T any] struct { //nolint:govet // OK
	collection *mongo.Collection
	filter     bsoncore.Document
	options    *options.FindOneOptions
}

func (x FindOperation[T]) Find(ctx context.Context) (T, error) {
	var document T

	if err := x.collection.FindOne(ctx, x.filter, x.options).Decode(&document); err != nil {
		return zero[T](), err
	}

	return document, nil
}

func Find[T any](c *mongo.Collection, opts ...FindOption[T]) FindOperation[T] {
	x := FindOperation[T]{
		collection: c,
		filter:     bsoncore.NewDocumentBuilder().Build(),
		options:    nil,
	}

	for _, opt := range opts {
		opt(&x)
	}

	return x
}

type FindOption[T any] func(x *FindOperation[T])

func WithFindFilter[T any](filter bsoncore.Document) FindOption[T] {
	return func(x *FindOperation[T]) {
		x.filter = filter
	}
}

func WithFindOptions[T any](opt *options.FindOneOptions) FindOption[T] {
	return func(x *FindOperation[T]) {
		x.options = opt
	}
}

func zero[T any]() T {
	return *new(T)
}
