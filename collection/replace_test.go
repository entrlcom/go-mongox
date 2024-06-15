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

var _ = ginkgo.Describe("Replace", func() {
	type input struct {
		document Document
		filter   bsoncore.Document
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

	document := Document{
		ID:   gofakeit.UUID(),
		Name: gofakeit.Name(),
	}

	ginkgo.JustBeforeEach(func(ctx ginkgo.SpecContext) {
		err := collection.Insert[Document](c).Insert(ctx, document)
		gomega.Expect(err).Should(gomega.Succeed())
	})

	document.Name = "REDACTED FOR PRIVACY"

	ginkgo.DescribeTable("", func(ctx context.Context, in input, out output) {
		err := collection.Replace[Document](
			c,
			collection.WithReplaceFilter[Document](in.filter),
		).Replace(ctx, in.document)
		gomega.Expect(err).Should(gomega.Succeed())

		v, err := collection.FindMany[Document](
			c,
			collection.WithFindManyFilter[Document](in.filter),
		).FindMany(ctx)

		switch out.err {
		case nil:
			gomega.Expect(err).Should(gomega.Succeed())

			for _, document := range v {
				gomega.Expect(document).To(gomega.Equal(in.document))
			}
		default:
			gomega.Expect(err).To(gomega.MatchError(out.err))
			gomega.Expect(v).To(gomega.BeZero())
		}
	},
		ginkgo.Entry("error (mongo.ErrNoDocuments)",
			input{document: document, filter: query.Eq().Field("id").Null()},
			output{err: mongo.ErrNoDocuments},
		),
		ginkgo.Entry("ok",
			input{
				document: document,
				filter:   query.Eq().Field("id").String(document.ID),
			},
			output{err: nil},
		),
	)
})
