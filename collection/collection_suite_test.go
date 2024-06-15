package collection_test

import (
	"testing"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

func Test(t *testing.T) {
	t.Parallel()

	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Collection Test Suite")
}

type Document struct {
	ID   string `bson:"id,omitempty"`
	Name string `bson:"name,omitempty"`
}
