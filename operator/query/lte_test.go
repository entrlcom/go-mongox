package query_test //nolint:dupl // OK._test //nolint:dupl // OK.

import (
	"time"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"

	"entrlcom.dev/mongox/operator/query"
)

var _ = ginkgo.Describe("Lte", func() {
	ginkgo.DescribeTable("", func(document bsoncore.Document, s string) {
		gomega.Expect(document.String()).To(gomega.Equal(s))
	},
		// bool.
		ginkgo.Entry("bool (true)",
			query.Lte().Field("k").Bool(true),
			`{"k": {"$lte": true}}`,
		),
		ginkgo.Entry("bool (false)",
			query.Lte().Field("k").Bool(false),
			`{"k": {"$lte": false}}`,
		),

		// float64.
		ginkgo.Entry("float64",
			query.Lte().Field("k").Float64(0),
			`{"k": {"$lte": {"$numberDouble":"0.0"}}}`,
		),

		// int32.
		ginkgo.Entry("int32",
			query.Lte().Field("k").Int32(0),
			`{"k": {"$lte": {"$numberInt":"0"}}}`,
		),

		// int64.
		ginkgo.Entry("int64",
			query.Lte().Field("k").Int64(0),
			`{"k": {"$lte": {"$numberLong":"0"}}}`,
		),

		// string.
		ginkgo.Entry("string",
			query.Lte().Field("k").String(""),
			`{"k": {"$lte": ""}}`,
		),

		// time.
		ginkgo.Entry("time",
			query.Lte().Field("k").Time(time.Unix(0, 0)),
			`{"k": {"$lte": {"$date":{"$numberLong":"0"}}}}`,
		),
	)
})
