package query_test //nolint:dupl // OK._test //nolint:dupl // OK.

import (
	"time"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"

	"entrlcom.dev/mongox/operator/query"
)

var _ = ginkgo.Describe("Eq", func() {
	ginkgo.DescribeTable("", func(document bsoncore.Document, s string) {
		gomega.Expect(document.String()).To(gomega.Equal(s))
	},
		// array.
		ginkgo.Entry("array (empty)",
			query.Eq().Field("k").Array(bsoncore.NewArrayBuilder().Build()),
			`{"k": {"$eq": []}}`,
		),
		ginkgo.Entry("array",
			query.Eq().Field("k").Array(bsoncore.NewArrayBuilder().
				AppendBoolean(true).
				AppendDouble(0.1).
				AppendNull().
				AppendString("John").
				Build(),
			),
			`{"k": {"$eq": [true,{"$numberDouble":"0.1"},null,"John"]}}`,
		),

		// bool.
		ginkgo.Entry("bool (true)",
			query.Eq().Field("k").Bool(true),
			`{"k": {"$eq": true}}`,
		),
		ginkgo.Entry("bool (false)",
			query.Eq().Field("k").Bool(false),
			`{"k": {"$eq": false}}`,
		),

		// document.
		ginkgo.Entry("document (empty)",
			query.Eq().Field("k").Document(bsoncore.NewDocumentBuilder().Build()),
			`{"k": {"$eq": {}}}`,
		),
		ginkgo.Entry("document",
			query.Eq().Field("k").Document(bsoncore.NewDocumentBuilder().AppendString("string", "string").Build()),
			`{"k": {"$eq": {"string": "string"}}}`,
		),

		// float64.
		ginkgo.Entry("float64",
			query.Eq().Field("k").Float64(0),
			`{"k": {"$eq": {"$numberDouble":"0.0"}}}`,
		),

		// int32.
		ginkgo.Entry("int32",
			query.Eq().Field("k").Int32(0),
			`{"k": {"$eq": {"$numberInt":"0"}}}`,
		),

		// int64.
		ginkgo.Entry("int64",
			query.Eq().Field("k").Int64(0),
			`{"k": {"$eq": {"$numberLong":"0"}}}`,
		),

		// null.
		ginkgo.Entry("null",
			query.Eq().Field("k").Null(),
			`{"k": {"$eq": null}}`,
		),

		// string.
		ginkgo.Entry("string",
			query.Eq().Field("k").String(""),
			`{"k": {"$eq": ""}}`,
		),

		// time.
		ginkgo.Entry("time",
			query.Eq().Field("k").Time(time.Unix(0, 0)),
			`{"k": {"$eq": {"$date":{"$numberLong":"0"}}}}`,
		),
	)
})
