package container_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "cosmos-go-sdk/container"
	"cosmos-go-sdk/testdata/mocks"
)

var _ = Describe("Container", func() {

	var (
		mockCtrl      *gomock.Controller
		mockRequest   *mocks.MockIRequest
		testContainer Container
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockRequest = mocks.NewMockIRequest(mockCtrl)
		testContainer = Container{
			Name:    "testName",
			DbName:  "testDBName",
			URI:     "testURI",
			Key:     "testKey",
			Request: mockRequest,
		}
	})

	Context("Client", func() {
		It("should successfully return a new instance of a Container Client", func() {
			Expect(testContainer).To(BeAssignableToTypeOf(Container{}))
		})
	})

	Context("Read", func() {
		It("should successfully read current instance of Container", func() {
			mockRequest.EXPECT().Get().Return([]byte("Successful Container Read"), nil).Times(1)
			container, testReadError := testContainer.Read()
			Expect(testReadError).To(BeNil())
			Expect(container).ToNot(BeNil())
		})
	})

	Context("Delete", func() {
		It("should successfully delete an Container", func() {
			testDeleteError := testContainer.Delete()
			Expect(testDeleteError).To(BeNil())
		})
	})

	Context("Replace", func() {
		It("should successfully replace an Container", func() {
			testEntity := Entity{}
			container, testReplaceError := testContainer.Replace(testEntity)
			Expect(testReplaceError).To(BeNil())
			Expect(container).ToNot(BeNil())
		})
	})
})
