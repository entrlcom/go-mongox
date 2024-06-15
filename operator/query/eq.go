package query //nolint:dupl // OK.

import (
	"time"

	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
)

type EqBuilder interface {
	Field(v string) EqFieldBuilder
}

type EqFieldBuilder interface {
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

type eqBuilder struct{}

func (x eqBuilder) Field(field string) EqFieldBuilder {
	return eqFieldBuilder{field: field}
}

func Eq() EqBuilder {
	return eqBuilder{}
}

type eqFieldBuilder struct {
	field string
}

func (x eqFieldBuilder) Array(array bsoncore.Array) bsoncore.Document {
	return bsoncore.NewDocumentBuilder().
		AppendDocument(x.field, bsoncore.NewDocumentBuilder().
			AppendArray("$eq", array).
			Build()).
		Build()
}

func (x eqFieldBuilder) Bool(v bool) bsoncore.Document {
	return bsoncore.NewDocumentBuilder().
		AppendDocument(x.field, bsoncore.NewDocumentBuilder().
			AppendBoolean("$eq", v).
			Build()).
		Build()
}

func (x eqFieldBuilder) Document(document bsoncore.Document) bsoncore.Document {
	return bsoncore.NewDocumentBuilder().
		AppendDocument(x.field, bsoncore.NewDocumentBuilder().
			AppendDocument("$eq", document).
			Build()).
		Build()
}

func (x eqFieldBuilder) Float64(v float64) bsoncore.Document {
	return bsoncore.NewDocumentBuilder().
		AppendDocument(x.field, bsoncore.NewDocumentBuilder().
			AppendDouble("$eq", v).
			Build()).
		Build()
}

func (x eqFieldBuilder) Int32(v int32) bsoncore.Document {
	return bsoncore.NewDocumentBuilder().
		AppendDocument(x.field, bsoncore.NewDocumentBuilder().
			AppendInt32("$eq", v).
			Build()).
		Build()
}

func (x eqFieldBuilder) Int64(v int64) bsoncore.Document {
	return bsoncore.NewDocumentBuilder().
		AppendDocument(x.field, bsoncore.NewDocumentBuilder().
			AppendInt64("$eq", v).
			Build()).
		Build()
}

func (x eqFieldBuilder) Null() bsoncore.Document {
	return bsoncore.NewDocumentBuilder().
		AppendDocument(x.field, bsoncore.NewDocumentBuilder().
			AppendNull("$eq").
			Build()).
		Build()
}

func (x eqFieldBuilder) String(v string) bsoncore.Document {
	return bsoncore.NewDocumentBuilder().
		AppendDocument(x.field, bsoncore.NewDocumentBuilder().
			AppendString("$eq", v).
			Build()).
		Build()
}

func (x eqFieldBuilder) Time(v time.Time) bsoncore.Document {
	return bsoncore.NewDocumentBuilder().
		AppendDocument(x.field, bsoncore.NewDocumentBuilder().
			AppendDateTime("$eq", v.UnixMilli()).
			Build()).
		Build()
}

var (
	_ EqBuilder      = (*eqBuilder)(nil)
	_ EqFieldBuilder = (*eqFieldBuilder)(nil)
)
