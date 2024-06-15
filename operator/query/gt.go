package query //nolint:dupl // OK.

import (
	"time"

	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
)

type GtBuilder interface {
	Field(v string) GtFieldBuilder
}

type GtFieldBuilder interface {
	Bool(v bool) bsoncore.Document
	Float64(v float64) bsoncore.Document
	Int32(v int32) bsoncore.Document
	Int64(v int64) bsoncore.Document
	String(v string) bsoncore.Document
	Time(v time.Time) bsoncore.Document
}

type gtBuilder struct{}

func (x gtBuilder) Field(field string) GtFieldBuilder {
	return gtFieldBuilder{field: field}
}

func Gt() GtBuilder {
	return gtBuilder{}
}

type gtFieldBuilder struct {
	field string
}

func (x gtFieldBuilder) Bool(v bool) bsoncore.Document {
	return bsoncore.NewDocumentBuilder().
		AppendDocument(x.field, bsoncore.NewDocumentBuilder().
			AppendBoolean("$gt", v).
			Build()).
		Build()
}

func (x gtFieldBuilder) Float64(v float64) bsoncore.Document {
	return bsoncore.NewDocumentBuilder().
		AppendDocument(x.field, bsoncore.NewDocumentBuilder().
			AppendDouble("$gt", v).
			Build()).
		Build()
}

func (x gtFieldBuilder) Int32(v int32) bsoncore.Document {
	return bsoncore.NewDocumentBuilder().
		AppendDocument(x.field, bsoncore.NewDocumentBuilder().
			AppendInt32("$gt", v).
			Build()).
		Build()
}

func (x gtFieldBuilder) Int64(v int64) bsoncore.Document {
	return bsoncore.NewDocumentBuilder().
		AppendDocument(x.field, bsoncore.NewDocumentBuilder().
			AppendInt64("$gt", v).
			Build()).
		Build()
}

func (x gtFieldBuilder) String(v string) bsoncore.Document {
	return bsoncore.NewDocumentBuilder().
		AppendDocument(x.field, bsoncore.NewDocumentBuilder().
			AppendString("$gt", v).
			Build()).
		Build()
}

func (x gtFieldBuilder) Time(v time.Time) bsoncore.Document {
	return bsoncore.NewDocumentBuilder().
		AppendDocument(x.field, bsoncore.NewDocumentBuilder().
			AppendDateTime("$gt", v.UnixMilli()).
			Build()).
		Build()
}

var (
	_ GtBuilder      = (*gtBuilder)(nil)
	_ GtFieldBuilder = (*gtFieldBuilder)(nil)
)
