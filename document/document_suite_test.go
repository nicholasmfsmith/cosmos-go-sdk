package document_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDocument(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Document Suite")
}
