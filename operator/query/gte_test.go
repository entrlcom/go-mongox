package query_test //nolint:dupl // OK.

import (
	"time"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"

	"entrlcom.dev/mongox/operator/query"
)

var _ = ginkgo.Describe("Gte", func() {
	ginkgo.DescribeTable("", func(document bsoncore.Document, s string) {
		gomega.Expect(document.String()).To(gomega.Equal(s))
	},
		// bool.
		ginkgo.Entry("bool (true)",
			query.Gte().Field("k").Bool(true),
			`{"k": {"$gte": true}}`,
		),
		ginkgo.Entry("bool (false)",
			query.Gte().Field("k").Bool(false),
			`{"k": {"$gte": false}}`,
		),

		// float64.
		ginkgo.Entry("float64",
			query.Gte().Field("k").Float64(0),
			`{"k": {"$gte": {"$numberDouble":"0.0"}}}`,
		),

		// int32.
		ginkgo.Entry("int32",
			query.Gte().Field("k").Int32(0),
			`{"k": {"$gte": {"$numberInt":"0"}}}`,
		),

		// int64.
		ginkgo.Entry("int64",
			query.Gte().Field("k").Int64(0),
			`{"k": {"$gte": {"$numberLong":"0"}}}`,
		),

		// string.
		ginkgo.Entry("string",
			query.Gte().Field("k").String(""),
			`{"k": {"$gte": ""}}`,
		),

		// time.
		ginkgo.Entry("time",
			query.Gte().Field("k").Time(time.Unix(0, 0)),
			`{"k": {"$gte": {"$date":{"$numberLong":"0"}}}}`,
		),
	)
})
