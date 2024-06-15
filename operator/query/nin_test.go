package query_test //nolint:dupl // OK._test //nolint:dupl // OK.

import (
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"

	"entrlcom.dev/mongox/operator/query"
)

var _ = ginkgo.Describe("Nin", func() {
	ginkgo.DescribeTable("", func(document bsoncore.Document, s string) {
		gomega.Expect(document.String()).To(gomega.Equal(s))
	},
		ginkgo.Entry("array (empty)",
			query.Nin().Field("name").Array(bsoncore.NewArrayBuilder().Build()),
			`{"name": {"$nin": []}}`,
		),
		ginkgo.Entry("array",
			query.Nin().Field("name").Array(bsoncore.NewArrayBuilder().
				AppendBoolean(true).
				AppendDouble(0.1).
				AppendNull().
				AppendString("John").
				Build(),
			),
			`{"name": {"$nin": [true,{"$numberDouble":"0.1"},null,"John"]}}`,
		),
	)
})
