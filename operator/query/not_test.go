package query_test

import (
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"

	"entrlcom.dev/mongox/operator/query"
)

var _ = ginkgo.Describe("Not", func() {
	ginkgo.DescribeTable("", func(document bsoncore.Document, s string) {
		gomega.Expect(document.String()).To(gomega.Equal(s))
	},
		ginkgo.Entry("not (empty)",
			query.Not(bsoncore.NewDocumentBuilder().Build()),
			`{"$not": {}}`,
		),
		ginkgo.Entry("not",
			query.Not(query.Eq().Field("k").Null()),
			`{"$not": {"k": {"$eq": null}}}`,
		),
	)
})
