package query

import (
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
)

func Not(document bsoncore.Document) bsoncore.Document {
	return bsoncore.NewDocumentBuilder().AppendDocument("$not", document).Build()
}
