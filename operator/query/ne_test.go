package query_test //nolint:dupl // OK._test //nolint:dupl // OK.

import (
	"time"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"

	"entrlcom.dev/mongox/operator/query"
)

var _ = ginkgo.Describe("Ne", func() {
	ginkgo.DescribeTable("", func(document bsoncore.Document, s string) {
		gomega.Expect(document.String()).To(gomega.Equal(s))
	},
		// array.
		ginkgo.Entry("array (empty)",
			query.Ne().Field("k").Array(bsoncore.NewArrayBuilder().Build()),
			`{"k": {"$ne": []}}`,
		),
		ginkgo.Entry("array",
			query.Ne().Field("k").Array(bsoncore.NewArrayBuilder().
				AppendBoolean(true).
				AppendDouble(0.1).
				AppendNull().
				AppendString("John").
				Build(),
			),
			`{"k": {"$ne": [true,{"$numberDouble":"0.1"},null,"John"]}}`,
		),

		// bool.
		ginkgo.Entry("bool (true)",
			query.Ne().Field("k").Bool(true),
			`{"k": {"$ne": true}}`,
		),
		ginkgo.Entry("bool (false)",
			query.Ne().Field("k").Bool(false),
			`{"k": {"$ne": false}}`,
		),

		// document.
		ginkgo.Entry("document (empty)",
			query.Ne().Field("k").Document(bsoncore.NewDocumentBuilder().Build()),
			`{"k": {"$ne": {}}}`,
		),
		ginkgo.Entry("document",
			query.Ne().Field("k").Document(bsoncore.NewDocumentBuilder().AppendString("string", "string").Build()),
			`{"k": {"$ne": {"string": "string"}}}`,
		),

		// float64.
		ginkgo.Entry("float64",
			query.Ne().Field("k").Float64(0),
			`{"k": {"$ne": {"$numberDouble":"0.0"}}}`,
		),

		// int32.
		ginkgo.Entry("int32",
			query.Ne().Field("k").Int32(0),
			`{"k": {"$ne": {"$numberInt":"0"}}}`,
		),

		// int64.
		ginkgo.Entry("int64",
			query.Ne().Field("k").Int64(0),
			`{"k": {"$ne": {"$numberLong":"0"}}}`,
		),

		// null.
		ginkgo.Entry("null",
			query.Ne().Field("k").Null(),
			`{"k": {"$ne": null}}`,
		),

		// string.
		ginkgo.Entry("string",
			query.Ne().Field("k").String(""),
			`{"k": {"$ne": ""}}`,
		),

		// time.
		ginkgo.Entry("time",
			query.Ne().Field("k").Time(time.Unix(0, 0)),
			`{"k": {"$ne": {"$date":{"$numberLong":"0"}}}}`,
		),
	)
})
