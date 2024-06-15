package collection_test

import (
	"context"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/mongo"

	"entrlcom.dev/mongox/collection"
	"entrlcom.dev/mongox/docker"
	"entrlcom.dev/mongox/operator/query"
)

var _ = ginkgo.Describe("Insert", func() {
	type input struct {
		document Document
	}

	type output struct {
		err error
	}

	var c *mongo.Collection

	ginkgo.BeforeEach(func(ctx ginkgo.SpecContext) {
		mongo, err := docker.NewMongo(ctx)
		gomega.Expect(err).Should(gomega.Succeed())

		c = mongo.GetClient().Database("test").Collection("test")
	})

	ginkgo.DescribeTable("", func(ctx context.Context, in input, out output) {
		err := collection.Insert[Document](c).Insert(ctx, in.document)

		switch out.err {
		case nil:
			gomega.Expect(err).Should(gomega.Succeed())
		default:
			gomega.Expect(err).To(gomega.MatchError(out.err))
		}

		filter := query.Eq().Field("id").String(in.document.ID)

		v, err := collection.Find[Document](
			c,
			collection.WithFindFilter[Document](filter),
		).Find(ctx)
		gomega.Expect(err).Should(gomega.Succeed())
		gomega.Expect(v).To(gomega.Equal(in.document))
	},
		ginkgo.Entry("ok",
			input{document: Document{ID: gofakeit.UUID(), Name: gofakeit.Name()}},
			output{err: nil},
		),
	)
})
