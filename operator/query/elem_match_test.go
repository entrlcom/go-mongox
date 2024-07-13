package query_test

import (
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"

	"entrlcom.dev/mongox/operator/query"
)

var _ = ginkgo.Describe("ElemMatch", func() {
	ginkgo.DescribeTable("", func(document bsoncore.Document, s string) {
		gomega.Expect(document.String()).To(gomega.Equal(s))
	},
		ginkgo.Entry("document (empty)",
			query.ElemMatch().Field("instock").Document(bsoncore.NewDocumentBuilder().Build()),
			`{"instock": {"$elemMatch": {}}}`,
		),
		ginkgo.Entry("document",
			query.ElemMatch().Field("instock").Document(bsoncore.NewDocumentBuilder().
				AppendInt32("qty", 5).
				AppendString("warehouse", "A").
				Build(),
			),
			`{"instock": {"$elemMatch": {"qty": {"$numberInt":"5"},"warehouse": "A"}}}`,
		),
	)
})
