package query //nolint:dupl // OK.

import (
	"time"

	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
)

type NeBuilder interface {
	Field(v string) NeFieldBuilder
}

type NeFieldBuilder interface {
	Array(v bsoncore.Array) bsoncore.Document
	Bool(v bool) bsoncore.Document
	Document(document bsoncore.Document) bsoncore.Document
	Float64(v float64) bsoncore.Document
	Int32(v int32) bsoncore.Document
	Int64(v int64) bsoncore.Document
	Null() bsoncore.Document
	String(v string) bsoncore.Document
	Time(v time.Time) bsoncore.Document
}

type neBuilder struct{}

func (x neBuilder) Field(field string) NeFieldBuilder {
	return neFieldBuilder{field: field}
}

func Ne() NeBuilder {
	return neBuilder{}
}

type neFieldBuilder struct {
	field string
}

func (x neFieldBuilder) Array(array bsoncore.Array) bsoncore.Document {
	return bsoncore.NewDocumentBuilder().
		AppendDocument(x.field, bsoncore.NewDocumentBuilder().
			AppendArray("$ne", array).
			Build()).
		Build()
}

func (x neFieldBuilder) Bool(v bool) bsoncore.Document {
	return bsoncore.NewDocumentBuilder().
		AppendDocument(x.field, bsoncore.NewDocumentBuilder().
			AppendBoolean("$ne", v).
			Build()).
		Build()
}

func (x neFieldBuilder) Document(document bsoncore.Document) bsoncore.Document {
	return bsoncore.NewDocumentBuilder().
		AppendDocument(x.field, bsoncore.NewDocumentBuilder().
			AppendDocument("$ne", document).
			Build()).
		Build()
}

func (x neFieldBuilder) Float64(v float64) bsoncore.Document {
	return bsoncore.NewDocumentBuilder().
		AppendDocument(x.field, bsoncore.NewDocumentBuilder().
			AppendDouble("$ne", v).
			Build()).
		Build()
}

func (x neFieldBuilder) Int32(v int32) bsoncore.Document {
	return bsoncore.NewDocumentBuilder().
		AppendDocument(x.field, bsoncore.NewDocumentBuilder().
			AppendInt32("$ne", v).
			Build()).
		Build()
}

func (x neFieldBuilder) Int64(v int64) bsoncore.Document {
	return bsoncore.NewDocumentBuilder().
		AppendDocument(x.field, bsoncore.NewDocumentBuilder().
			AppendInt64("$ne", v).
			Build()).
		Build()
}

func (x neFieldBuilder) Null() bsoncore.Document {
	return bsoncore.NewDocumentBuilder().
		AppendDocument(x.field, bsoncore.NewDocumentBuilder().
			AppendNull("$ne").
			Build()).
		Build()
}

func (x neFieldBuilder) String(v string) bsoncore.Document {
	return bsoncore.NewDocumentBuilder().
		AppendDocument(x.field, bsoncore.NewDocumentBuilder().
			AppendString("$ne", v).
			Build()).
		Build()
}

func (x neFieldBuilder) Time(v time.Time) bsoncore.Document {
	return bsoncore.NewDocumentBuilder().
		AppendDocument(x.field, bsoncore.NewDocumentBuilder().
			AppendDateTime("$ne", v.UnixMilli()).
			Build()).
		Build()
}

var (
	_ NeBuilder      = (*neBuilder)(nil)
	_ NeFieldBuilder = (*neFieldBuilder)(nil)
)
