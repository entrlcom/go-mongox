package query //nolint:dupl // OK.

import (
	"time"

	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
)

type LtBuilder interface {
	Field(v string) LtFieldBuilder
}

type LtFieldBuilder interface {
	Bool(v bool) bsoncore.Document
	Float64(v float64) bsoncore.Document
	Int32(v int32) bsoncore.Document
	Int64(v int64) bsoncore.Document
	String(v string) bsoncore.Document
	Time(v time.Time) bsoncore.Document
}

type ltBuilder struct{}

func (x ltBuilder) Field(field string) LtFieldBuilder {
	return ltFieldBuilder{field: field}
}

func Lt() LtBuilder {
	return ltBuilder{}
}

type ltFieldBuilder struct {
	field string
}

func (x ltFieldBuilder) Bool(v bool) bsoncore.Document {
	return bsoncore.NewDocumentBuilder().
		AppendDocument(x.field, bsoncore.NewDocumentBuilder().
			AppendBoolean("$lt", v).
			Build()).
		Build()
}

func (x ltFieldBuilder) Float64(v float64) bsoncore.Document {
	return bsoncore.NewDocumentBuilder().
		AppendDocument(x.field, bsoncore.NewDocumentBuilder().
			AppendDouble("$lt", v).
			Build()).
		Build()
}

func (x ltFieldBuilder) Int32(v int32) bsoncore.Document {
	return bsoncore.NewDocumentBuilder().
		AppendDocument(x.field, bsoncore.NewDocumentBuilder().
			AppendInt32("$lt", v).
			Build()).
		Build()
}

func (x ltFieldBuilder) Int64(v int64) bsoncore.Document {
	return bsoncore.NewDocumentBuilder().
		AppendDocument(x.field, bsoncore.NewDocumentBuilder().
			AppendInt64("$lt", v).
			Build()).
		Build()
}

func (x ltFieldBuilder) String(v string) bsoncore.Document {
	return bsoncore.NewDocumentBuilder().
		AppendDocument(x.field, bsoncore.NewDocumentBuilder().
			AppendString("$lt", v).
			Build()).
		Build()
}

func (x ltFieldBuilder) Time(v time.Time) bsoncore.Document {
	return bsoncore.NewDocumentBuilder().
		AppendDocument(x.field, bsoncore.NewDocumentBuilder().
			AppendDateTime("$lt", v.UnixMilli()).
			Build()).
		Build()
}

var (
	_ LtBuilder      = (*ltBuilder)(nil)
	_ LtFieldBuilder = (*ltFieldBuilder)(nil)
)
