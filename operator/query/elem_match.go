package query

import (
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
)

type ElemMatchBuilder interface {
	Field(v string) ElemMatchFieldBuilder
}

type ElemMatchFieldBuilder interface {
	Document(v bsoncore.Document) bsoncore.Document
}

type elemMatchBuilder struct{}

func (x elemMatchBuilder) Field(field string) ElemMatchFieldBuilder {
	return elemMatchFieldBuilder{field: field}
}

func ElemMatch() ElemMatchBuilder {
	return elemMatchBuilder{}
}

type elemMatchFieldBuilder struct {
	field string
}

func (x elemMatchFieldBuilder) Document(document bsoncore.Document) bsoncore.Document {
	return bsoncore.NewDocumentBuilder().
		AppendDocument(x.field, bsoncore.NewDocumentBuilder().
			AppendDocument("$elemMatch", document).
			Build()).
		Build()
}
