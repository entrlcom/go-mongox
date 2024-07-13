package query_test //nolint:dupl // OK.

import (
	"time"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"

	"entrlcom.dev/mongox/operator/query"
)

var _ = ginkgo.Describe("Gt", func() {
	ginkgo.DescribeTable("", func(document bsoncore.Document, s string) {
		gomega.Expect(document.String()).To(gomega.Equal(s))
	},
		// bool.
		ginkgo.Entry("bool (true)",
			query.Gt().Field("k").Bool(true),
			`{"k": {"$gt": true}}`,
		),
		ginkgo.Entry("bool (false)",
			query.Gt().Field("k").Bool(false),
			`{"k": {"$gt": false}}`,
		),

		// float64.
		ginkgo.Entry("float64",
			query.Gt().Field("k").Float64(0),
			`{"k": {"$gt": {"$numberDouble":"0.0"}}}`,
		),

		// int32.
		ginkgo.Entry("int32",
			query.Gt().Field("k").Int32(0),
			`{"k": {"$gt": {"$numberInt":"0"}}}`,
		),

		// int64.
		ginkgo.Entry("int64",
			query.Gt().Field("k").Int64(0),
			`{"k": {"$gt": {"$numberLong":"0"}}}`,
		),

		// string.
		ginkgo.Entry("string",
			query.Gt().Field("k").String(""),
			`{"k": {"$gt": ""}}`,
		),

		// time.
		ginkgo.Entry("time",
			query.Gt().Field("k").Time(time.Unix(0, 0)),
			`{"k": {"$gt": {"$date":{"$numberLong":"0"}}}}`,
		),
	)
})
