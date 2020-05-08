package document_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "cosmos-go-sdk/document"
	"cosmos-go-sdk/testdata/mocks"
)

var _ = Describe("document", func() {
	var (
		mockCtrl     *gomock.Controller
		mockRequest  *mocks.MockIRequest
		testdocument Document
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockRequest = mocks.NewMockIRequest(mockCtrl)
		testdocument = Document{
			ID:           "testID",
			PartitionKey: "testPartitionKey",
			URI:          "testURI",
			Key:          "testKey",
			Request:      mockRequest,
		}
	})

	Context("New", func() {
		It("should successfully return a new instance of Document", func() {
			testdocument = New("testID", "testPartitionKey", "containerURI", "testKey")
			Expect(testdocument).To(Not(BeNil()))
			Expect(testdocument).To(BeAssignableToTypeOf(Document{}))
		})
		It("should successfully return a new instance of Document with the correct ID", func() {
			testdocument = New("testID", "testPartitionKey", "containerURI", "testKey")
			Expect(testdocument.ID).To(Equal("testID"))
		})
		It("should successfully return a new instance of Document with the correct PartitionKey", func() {
			testdocument = New("testID", "testPartitionKey", "containerURI", "testKey")
			Expect(testdocument.PartitionKey).To(Equal("testPartitionKey"))
		})
		It("should successfully return a new instance of Document with the correct URI", func() {
			testdocument = New("testID", "testPartitionKey", "containerURI", "testKey")
			Expect(testdocument.URI).To(Equal("containerURI" + "/docs/testID"))
		})
		It("should successfully return a new instance of Document with the correct Key", func() {
			testdocument = New("testID", "testPartitionKey", "containerURI", "testKey")
			Expect(testdocument.Key).To(Equal("testKey"))
		})
	})

	Context("Create", func() {
		It("should successfully create a document in the Azure Cosmos Database Container", func() {
			testDocument := []byte("This is a test created document")
			mockRequest.EXPECT().Post(testDocument).Return([]byte("This is a test created document that was created"), nil).Times(1)
			createdDocument, testCreateError := testdocument.Create(testDocument)
			Expect(testCreateError).To(BeNil())
			Expect(createdDocument).To(Equal([]byte("This is a test created document that was created")))
		})
	})

	Context("Read", func() {
		It("should successfully read a document from the Azure Cosmos Database Container", func() {
			mockRequest.EXPECT().Get().Return([]byte("test document has been read"), nil).Times(1)
			document, testReadError := testdocument.Read()
			Expect(testReadError).To(BeNil())
			Expect(document).To(Equal([]byte("test document has been read")))
		})
	})

	Context("Update", func() {
		It("should successfully update a document in the Azure Cosmos Database Container", func() {
			testDocument := []byte("This is a test updated document")
			testUpdateError := testdocument.Update(testDocument)
			Expect(testUpdateError).To(BeNil())
		})
	})

	Context("Delete", func() {
		It("should successfully delete a document from the Azure Cosmos Database Container", func() {
			testDeleteError := testdocument.Delete()
			Expect(testDeleteError).To(BeNil())
		})
	})
})
