package query //nolint:dupl // OK.

import (
	"time"

	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
)

type LteBuilder interface {
	Field(v string) LteFieldBuilder
}

type LteFieldBuilder interface {
	Bool(v bool) bsoncore.Document
	Float64(v float64) bsoncore.Document
	Int32(v int32) bsoncore.Document
	Int64(v int64) bsoncore.Document
	String(v string) bsoncore.Document
	Time(v time.Time) bsoncore.Document
}

type lteBuilder struct{}

func (x lteBuilder) Field(field string) LteFieldBuilder {
	return lteFieldBuilder{field: field}
}

func Lte() LteBuilder {
	return lteBuilder{}
}

type lteFieldBuilder struct {
	field string
}

func (x lteFieldBuilder) Bool(v bool) bsoncore.Document {
	return bsoncore.NewDocumentBuilder().
		AppendDocument(x.field, bsoncore.NewDocumentBuilder().
			AppendBoolean("$lte", v).
			Build()).
		Build()
}

func (x lteFieldBuilder) Float64(v float64) bsoncore.Document {
	return bsoncore.NewDocumentBuilder().
		AppendDocument(x.field, bsoncore.NewDocumentBuilder().
			AppendDouble("$lte", v).
			Build()).
		Build()
}

func (x lteFieldBuilder) Int32(v int32) bsoncore.Document {
	return bsoncore.NewDocumentBuilder().
		AppendDocument(x.field, bsoncore.NewDocumentBuilder().
			AppendInt32("$lte", v).
			Build()).
		Build()
}

func (x lteFieldBuilder) Int64(v int64) bsoncore.Document {
	return bsoncore.NewDocumentBuilder().
		AppendDocument(x.field, bsoncore.NewDocumentBuilder().
			AppendInt64("$lte", v).
			Build()).
		Build()
}

func (x lteFieldBuilder) String(v string) bsoncore.Document {
	return bsoncore.NewDocumentBuilder().
		AppendDocument(x.field, bsoncore.NewDocumentBuilder().
			AppendString("$lte", v).
			Build()).
		Build()
}

func (x lteFieldBuilder) Time(v time.Time) bsoncore.Document {
	return bsoncore.NewDocumentBuilder().
		AppendDocument(x.field, bsoncore.NewDocumentBuilder().
			AppendDateTime("$lte", v.UnixMilli()).
			Build()).
		Build()
}

var (
	_ LteBuilder      = (*lteBuilder)(nil)
	_ LteFieldBuilder = (*lteFieldBuilder)(nil)
)
