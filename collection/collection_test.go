package collection_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "cosmos-go-sdk/collection"
	libdocument "cosmos-go-sdk/document"
	"cosmos-go-sdk/testdata/mocks"
)

var _ = Describe("Collection", func() {

	var (
		mockCtrl       *gomock.Controller
		mockRequest    *mocks.MockIRequest
		testCollection Collection
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockRequest = mocks.NewMockIRequest(mockCtrl)
		testCollection = Collection{
			Name:    "testName",
			URI:     "testURI",
			Key:     "testKey",
			Request: mockRequest,
		}
	})

	Context("New", func() {
		It("should successfully return a new instance of a Collection", func() {
			testCollection = New("testName", "testDatabaseURI", "testKey")
			Expect(testCollection).To(BeAssignableToTypeOf(Collection{}))
		})
	})

	Context("Read", func() {
		It("should successfully read current instance of Collection", func() {
			mockRequest.EXPECT().Get().Return([]byte("Successful Collection Read"), nil).Times(1)
			collection, testReadError := testCollection.Read()
			Expect(testReadError).To(BeNil())
			Expect(collection).ToNot(BeNil())
		})
	})

	Context("Delete", func() {
		It("should successfully delete a Collection", func() {
			testDeleteError := testCollection.Delete()
			Expect(testDeleteError).To(BeNil())
		})
	})

	Context("Replace", func() {
		It("should successfully replace a Collection", func() {
			testEntity := Entity{}
			collection, testReplaceError := testCollection.Replace(testEntity)
			Expect(testReplaceError).To(BeNil())
			Expect(collection).ToNot(BeNil())
		})
	})

	Context("Document", func() {
		It("should successfully return a new instance of Document of current Collection", func() {
			document := testCollection.Document("testID")
			Expect(document).To(BeAssignableToTypeOf(libdocument.Document{}))
		})

	})
})
