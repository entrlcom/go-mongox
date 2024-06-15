package collection

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
)

type ReplaceOperation[T any] struct { //nolint:govet // OK.
	collection *mongo.Collection
	filter     bsoncore.Document
	options    *options.ReplaceOptions
}

func (x ReplaceOperation[T]) Replace(ctx context.Context, document T) error {
	_, err := x.collection.ReplaceOne(ctx, x.filter, document, x.options)

	return err
}

func Replace[T any](c *mongo.Collection, opts ...ReplaceOption[T]) ReplaceOperation[T] {
	x := ReplaceOperation[T]{
		collection: c,
		filter:     bsoncore.NewDocumentBuilder().Build(),
		options:    nil,
	}

	for _, opt := range opts {
		opt(&x)
	}

	return x
}

type ReplaceOption[T any] func(x *ReplaceOperation[T])

func WithReplaceFilter[T any](filter bsoncore.Document) ReplaceOption[T] {
	return func(x *ReplaceOperation[T]) {
		x.filter = filter
	}
}

func WithReplaceOptions[T any](opt *options.ReplaceOptions) ReplaceOption[T] {
	return func(x *ReplaceOperation[T]) {
		x.options = opt
	}
}
