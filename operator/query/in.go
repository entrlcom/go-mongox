package query //nolint:dupl // OK.

import (
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
)

type InBuilder interface {
	Field(v string) InFieldBuilder
}

type InFieldBuilder interface {
	Array(v bsoncore.Array) bsoncore.Document
}

type inBuilder struct{}

func (x inBuilder) Field(field string) InFieldBuilder {
	return inFieldBuilder{field: field}
}

func In() InBuilder {
	return inBuilder{}
}

type inFieldBuilder struct {
	field string
}

func (x inFieldBuilder) Array(array bsoncore.Array) bsoncore.Document {
	return bsoncore.NewDocumentBuilder().
		AppendDocument(x.field, bsoncore.NewDocumentBuilder().
			AppendArray("$in", array).
			Build()).
		Build()
}

var (
	_ InBuilder      = (*inBuilder)(nil)
	_ InFieldBuilder = (*inFieldBuilder)(nil)
)
