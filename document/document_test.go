package document_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "cosmos-go-sdk/document"
)

var _ = Describe("document", func() {
	var (
		testdocument Document
	)

	BeforeEach(func() {
		testdocument = Document{}
	})

	Context("New", func() {
		It("should successfully return a new instance of document", func() {
			testdocument = New("testId", "testPartitionKey", "containerURI", "testKey")
			Expect(testdocument).To(BeAssignableToTypeOf(Document{}))
		})
	})

	Context("Create", func() {
		It("should successfully create an document in the Azure Cosmos Database Container", func() {
			testDocument := []byte("This is a test new document")
			testCreateError := testdocument.Create(testDocument)
			Expect(testCreateError).To(BeNil())
		})
	})

	Context("Read", func() {
		It("should successfully read an document from the Azure Cosmos Database Container", func() {
			document, testReadError := testdocument.Read()
			Expect(testReadError).To(BeNil())
			Expect(document).ToNot(BeNil())
		})
	})

	Context("Update", func() {
		It("should successfully update an document in the Azure Cosmos Database Container", func() {
			testDocument := []byte("This is a test updated document")
			testUpdateError := testdocument.Update(testDocument)
			Expect(testUpdateError).To(BeNil())
		})
	})

	Context("Delete", func() {
		It("should successfully delete an document from the Azure Cosmos Database Container", func() {
			testDeleteError := testdocument.Delete()
			Expect(testDeleteError).To(BeNil())
		})
	})
})
