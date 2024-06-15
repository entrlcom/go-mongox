package collection

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
)

type CountOperation struct { //nolint:govet // OK
	approximately bool
	collection    *mongo.Collection
	filter        bsoncore.Document
	options       *options.CountOptions
}

func (x CountOperation) Count(ctx context.Context) (int64, error) {
	if x.approximately {
		return x.collection.EstimatedDocumentCount(ctx)
	}

	return x.collection.CountDocuments(ctx, x.filter, x.options)
}

func Count(c *mongo.Collection, opts ...CountOption) CountOperation {
	x := CountOperation{
		approximately: false,
		collection:    c,
		filter:        bsoncore.NewDocumentBuilder().Build(),
		options:       nil,
	}

	for _, opt := range opts {
		opt(&x)
	}

	return x
}

type CountOption func(x *CountOperation)

func WithCountApproximation() CountOption {
	return func(x *CountOperation) {
		x.approximately = true
	}
}

func WithCountFilter(filter bsoncore.Document) CountOption {
	return func(x *CountOperation) {
		x.filter = filter
	}
}

func WithCountOptions(opt *options.CountOptions) CountOption {
	return func(x *CountOperation) {
		x.options = opt
	}
}
