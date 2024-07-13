package query_test //nolint:dupl // OK.

import (
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"

	"entrlcom.dev/mongox/operator/query"
)

var _ = ginkgo.Describe("Or", func() {
	ginkgo.DescribeTable("", func(document bsoncore.Document, s string) {
		gomega.Expect(document.String()).To(gomega.Equal(s))
	},
		// 0.
		ginkgo.Entry("0",
			query.Or(),
			`{}`,
		),

		// 1.
		ginkgo.Entry("1",
			query.Or(query.Eq().Field("k").Null()),
			`{"k": {"$eq": null}}`,
		),

		// 2+.
		ginkgo.Entry("2+",
			query.Or(
				query.Eq().Field("k").Null(),
				query.Eq().Field("k").Null(),
			),
			`{"$or": [{"k": {"$eq": null}},{"k": {"$eq": null}}]}`,
		),
	)
})
