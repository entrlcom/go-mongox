package query_test //nolint:dupl // OK.

import (
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"

	"entrlcom.dev/mongox/operator/query"
)

var _ = ginkgo.Describe("In", func() {
	ginkgo.DescribeTable("", func(document bsoncore.Document, s string) {
		gomega.Expect(document.String()).To(gomega.Equal(s))
	},
		ginkgo.Entry("array (empty)",
			query.In().Field("name").Array(bsoncore.NewArrayBuilder().Build()),
			`{"name": {"$in": []}}`,
		),
		ginkgo.Entry("array",
			query.In().Field("name").Array(bsoncore.NewArrayBuilder().
				AppendBoolean(true).
				AppendDouble(0.1).
				AppendNull().
				AppendString("John").
				Build(),
			),
			`{"name": {"$in": [true,{"$numberDouble":"0.1"},null,"John"]}}`,
		),
	)
})
