package query_test //nolint:dupl // OK._test //nolint:dupl // OK.

import (
	"time"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"

	"entrlcom.dev/mongox/operator/query"
)

var _ = ginkgo.Describe("Lt", func() {
	ginkgo.DescribeTable("", func(document bsoncore.Document, s string) {
		gomega.Expect(document.String()).To(gomega.Equal(s))
	},
		// bool.
		ginkgo.Entry("bool (true)",
			query.Lt().Field("k").Bool(true),
			`{"k": {"$lt": true}}`,
		),
		ginkgo.Entry("bool (false)",
			query.Lt().Field("k").Bool(false),
			`{"k": {"$lt": false}}`,
		),

		// float64.
		ginkgo.Entry("float64",
			query.Lt().Field("k").Float64(0),
			`{"k": {"$lt": {"$numberDouble":"0.0"}}}`,
		),

		// int32.
		ginkgo.Entry("int32",
			query.Lt().Field("k").Int32(0),
			`{"k": {"$lt": {"$numberInt":"0"}}}`,
		),

		// int64.
		ginkgo.Entry("int64",
			query.Lt().Field("k").Int64(0),
			`{"k": {"$lt": {"$numberLong":"0"}}}`,
		),

		// string.
		ginkgo.Entry("string",
			query.Lt().Field("k").String(""),
			`{"k": {"$lt": ""}}`,
		),

		// time.
		ginkgo.Entry("time",
			query.Lt().Field("k").Time(time.Unix(0, 0)),
			`{"k": {"$lt": {"$date":{"$numberLong":"0"}}}}`,
		),
	)
})
