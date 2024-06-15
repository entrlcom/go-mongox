package collection_test

import (
	"context"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"

	"entrlcom.dev/mongox/collection"
	"entrlcom.dev/mongox/docker"
	"entrlcom.dev/mongox/operator/query"
)

var _ = ginkgo.Describe("Find", func() {
	type input struct {
		filter bsoncore.Document
	}

	type output struct { //nolint:govet // OK.
		document Document
		err      error
	}

	var c *mongo.Collection

	ginkgo.BeforeEach(func(ctx ginkgo.SpecContext) {
		mongo, err := docker.NewMongo(ctx)
		gomega.Expect(err).Should(gomega.Succeed())

		c = mongo.GetClient().Database("test").Collection("test")
	})

	document := Document{
		ID:   gofakeit.UUID(),
		Name: gofakeit.Name(),
	}

	ginkgo.JustBeforeEach(func(ctx ginkgo.SpecContext) {
		err := collection.Insert[Document](c).Insert(ctx, document)
		gomega.Expect(err).Should(gomega.Succeed())
	})

	ginkgo.DescribeTable("", func(ctx context.Context, in input, out output) {
		v, err := collection.Find[Document](
			c,
			collection.WithFindFilter[Document](in.filter),
		).Find(ctx)

		switch out.err {
		case nil:
			gomega.Expect(err).Should(gomega.Succeed())
			gomega.Expect(v).To(gomega.Equal(out.document))
		default:
			gomega.Expect(err).To(gomega.MatchError(out.err))
			gomega.Expect(v).To(gomega.BeZero())
		}
	},
		ginkgo.Entry("error (mongo.ErrNoDocuments)",
			input{filter: query.Eq().Field("id").Null()},
			output{document: document, err: mongo.ErrNoDocuments},
		),
		ginkgo.Entry("ok",
			input{filter: query.Eq().Field("id").String(document.ID)},
			output{document: document, err: nil},
		),
	)
})
