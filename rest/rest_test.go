package rest_test

import (
	"cosmos-go-sdk/mocks"
	. "cosmos-go-sdk/rest"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Rest", func() {
	// var resource IResource
	// var key string
	var (
		mockCtrl     *gomock.Controller
		mockResource *mocks.MockIResource
		headers      map[string]string
		key          string
	)

	// TODO: Dynamically test different resource types
	BeforeEach(func() {
		// resource = Database{
		// }
		mockCtrl = gomock.NewController(GinkgoT())
		mockResource = mocks.NewMockIResource(mockCtrl)
		headers = map[string]string{
			Authorization:   "testAuthorization",
			ContentType:     "testContentType",
			XMsDate:         "testXMsDate",
			XMsSessionToken: "testXMsSessionToken",
			XMsVersion:      "testXMsVersion",
		}
		key = "testKey"
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	// Context("Post", func() {
	// 	It("should successfully POST a resource in Azure", func() {
	// 		testPostResource, testPostError := Post(resource)
	// 		Expect(testPostResource).To(Not(BeNil()))
	// 		Expect(testPostError).To(BeNil())
	// 	})
	// })

	Context("Get", func() {
		It("should successfully GET a resource from Azure", func() {
			testGetResource, testGetError := Get(mockResource, headers, key)
			mockResource.EXPECT().URI().Return("", nil).Times(1)
			Expect(testGetResource).To(Not(BeNil()))
			Expect(testGetError).To(BeNil())
		})
	})

	// Context("Put", func() {
	// 	It("should successfully PUT a resource in Azure", func() {
	// 		testPutResource, testPutError := Put(resourceID, resource)
	// 		Expect(testPutResource).To(Not(BeNil()))
	// 		Expect(testPutError).To(BeNil())
	// 	})
	// })

	// Context("Delete", func() {
	// 	It("should successfully DELETE a resource in Azure", func() {
	// 		testDeleteError := Delete(resourceID)
	// 		Expect(testDeleteError).To(BeNil())
	// 	})
	// })
})
