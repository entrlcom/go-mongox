package query //nolint:dupl // OK.

import (
	"time"

	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
)

type GteBuilder interface {
	Field(v string) GteFieldBuilder
}

type GteFieldBuilder interface {
	Bool(v bool) bsoncore.Document
	Float64(v float64) bsoncore.Document
	Int32(v int32) bsoncore.Document
	Int64(v int64) bsoncore.Document
	String(v string) bsoncore.Document
	Time(v time.Time) bsoncore.Document
}

type gteBuilder struct{}

func (x gteBuilder) Field(field string) GteFieldBuilder {
	return gteFieldBuilder{field: field}
}

func Gte() GteBuilder {
	return gteBuilder{}
}

type gteFieldBuilder struct {
	field string
}

func (x gteFieldBuilder) Bool(v bool) bsoncore.Document {
	return bsoncore.NewDocumentBuilder().
		AppendDocument(x.field, bsoncore.NewDocumentBuilder().
			AppendBoolean("$gte", v).
			Build()).
		Build()
}

func (x gteFieldBuilder) Float64(v float64) bsoncore.Document {
	return bsoncore.NewDocumentBuilder().
		AppendDocument(x.field, bsoncore.NewDocumentBuilder().
			AppendDouble("$gte", v).
			Build()).
		Build()
}

func (x gteFieldBuilder) Int32(v int32) bsoncore.Document {
	return bsoncore.NewDocumentBuilder().
		AppendDocument(x.field, bsoncore.NewDocumentBuilder().
			AppendInt32("$gte", v).
			Build()).
		Build()
}

func (x gteFieldBuilder) Int64(v int64) bsoncore.Document {
	return bsoncore.NewDocumentBuilder().
		AppendDocument(x.field, bsoncore.NewDocumentBuilder().
			AppendInt64("$gte", v).
			Build()).
		Build()
}

func (x gteFieldBuilder) String(v string) bsoncore.Document {
	return bsoncore.NewDocumentBuilder().
		AppendDocument(x.field, bsoncore.NewDocumentBuilder().
			AppendString("$gte", v).
			Build()).
		Build()
}

func (x gteFieldBuilder) Time(v time.Time) bsoncore.Document {
	return bsoncore.NewDocumentBuilder().
		AppendDocument(x.field, bsoncore.NewDocumentBuilder().
			AppendDateTime("$gte", v.UnixMilli()).
			Build()).
		Build()
}

var (
	_ GteBuilder      = (*gteBuilder)(nil)
	_ GteFieldBuilder = (*gteFieldBuilder)(nil)
)
