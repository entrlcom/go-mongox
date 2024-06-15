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

var _ = ginkgo.Describe("Delete", func() {
	type input struct {
		filter bsoncore.Document
	}

	type output struct { //nolint:govet // OK.
		documents []Document
		err       error
	}

	var c *mongo.Collection

	ginkgo.BeforeEach(func(ctx ginkgo.SpecContext) {
		mongo, err := docker.NewMongo(ctx)
		gomega.Expect(err).Should(gomega.Succeed())

		c = mongo.GetClient().Database("test").Collection("test")
	})

	documents := make([]Document, 0)

	for i := 0; i < 2; i++ {
		documents = append(documents, Document{
			ID:   gofakeit.UUID(),
			Name: gofakeit.Name(),
		})
	}

	ginkgo.JustBeforeEach(func(ctx ginkgo.SpecContext) {
		for _, document := range documents {
			err := collection.Insert[Document](c).Insert(ctx, document)
			gomega.Expect(err).Should(gomega.Succeed())
		}
	})

	ginkgo.DescribeTable("", func(ctx context.Context, in input, out output) {
		err := collection.Delete[Document](
			c,
			collection.WithDeleteFilter[Document](in.filter),
		).Delete(ctx)

		switch out.err {
		case nil:
			gomega.Expect(err).Should(gomega.Succeed())
		default:
			gomega.Expect(err).To(gomega.MatchError(out.err))
		}

		count, err := collection.Count(
			c,
			collection.WithCountFilter(in.filter),
		).Count(ctx)
		gomega.Expect(err).Should(gomega.Succeed())
		gomega.Expect(count).Should(gomega.BeZero())
	},
		ginkgo.Entry("ok (1)",
			input{filter: query.Eq().Field("id").String(documents[0].ID)},
			output{documents: documents[:1], err: nil},
		),
		ginkgo.Entry("ok (2)",
			input{filter: query.In().Field("id").Array(bsoncore.NewArrayBuilder().
				AppendString(documents[0].ID).
				AppendString(documents[1].ID).
				Build())},
			output{documents: documents[:2], err: nil},
		),
	)
})
