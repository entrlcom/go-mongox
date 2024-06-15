package query

import (
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
)

func Or(documents ...bsoncore.Document) bsoncore.Document {
	switch len(documents) {
	case 0:
		return bsoncore.NewDocumentBuilder().Build()
	case 1:
		return documents[0]
	default:
		b := bsoncore.NewArrayBuilder()

		for _, document := range documents {
			b.AppendDocument(document)
		}

		return bsoncore.NewDocumentBuilder().AppendArray("$or", b.Build()).Build()
	}
}
