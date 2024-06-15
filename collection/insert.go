package collection

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type InsertOperation[T any] struct {
	collection *mongo.Collection
	options    *options.InsertOneOptions
}

func (x InsertOperation[T]) Insert(ctx context.Context, document T) error {
	_, err := x.collection.InsertOne(ctx, document, x.options)

	return err
}

func Insert[T any](c *mongo.Collection, opts ...InsertOption[T]) InsertOperation[T] {
	x := InsertOperation[T]{
		collection: c,
		options:    nil,
	}

	for _, opt := range opts {
		opt(&x)
	}

	return x
}

type InsertOption[T any] func(x *InsertOperation[T])

func WithInsertOptions[T any](opt *options.InsertOneOptions) InsertOption[T] {
	return func(x *InsertOperation[T]) {
		x.options = opt
	}
}
