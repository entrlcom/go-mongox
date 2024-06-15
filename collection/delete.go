package collection

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
)

type DeleteOperation[T any] struct { //nolint:govet // OK.
	collection *mongo.Collection
	filter     bsoncore.Document
	options    *options.DeleteOptions
}

func (x DeleteOperation[T]) Delete(ctx context.Context) error {
	_, err := x.collection.DeleteMany(ctx, x.filter, x.options)

	return err
}

func Delete[T any](c *mongo.Collection, opts ...DeleteOption[T]) DeleteOperation[T] {
	x := DeleteOperation[T]{
		collection: c,
		filter:     bsoncore.NewDocumentBuilder().Build(),
		options:    nil,
	}

	for _, opt := range opts {
		opt(&x)
	}

	return x
}

type DeleteOption[T any] func(x *DeleteOperation[T])

func WithDeleteFilter[T any](filter bsoncore.Document) DeleteOption[T] {
	return func(x *DeleteOperation[T]) {
		x.filter = filter
	}
}

func WithDeleteOptions[T any](opt *options.DeleteOptions) DeleteOption[T] {
	return func(x *DeleteOperation[T]) {
		x.options = opt
	}
}
