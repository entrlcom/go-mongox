package query //nolint:dupl // OK.

import (
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
)

type NinBuilder interface {
	Field(v string) NinFieldBuilder
}

type NinFieldBuilder interface {
	Array(v bsoncore.Array) bsoncore.Document
}

type ninBuilder struct{}

func (x ninBuilder) Field(field string) NinFieldBuilder {
	return ninFieldBuilder{field: field}
}

func Nin() NinBuilder {
	return ninBuilder{}
}

type ninFieldBuilder struct {
	field string
}

func (x ninFieldBuilder) Array(array bsoncore.Array) bsoncore.Document {
	return bsoncore.NewDocumentBuilder().
		AppendDocument(x.field, bsoncore.NewDocumentBuilder().
			AppendArray("$nin", array).
			Build()).
		Build()
}

var (
	_ NinBuilder      = (*ninBuilder)(nil)
	_ NinFieldBuilder = (*ninFieldBuilder)(nil)
)
