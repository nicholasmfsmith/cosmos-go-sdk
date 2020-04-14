package rest_test

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"cosmos-go-sdk/mocks"
	. "cosmos-go-sdk/rest"
)

var _ = Describe("Rest", func() {
	var (
		mockCtrl       *gomock.Controller
		mockResource   *mocks.MockIResource
		mockHttpClient *mocks.MockIHttpClient
		body           []byte
		id             string
		partitionKey   string
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockResource = mocks.NewMockIResource(mockCtrl)
		mockHttpClient = mocks.NewMockIHttpClient(mockCtrl)
		partitionKey = "testPartitionKey"
		body = []byte(`{"id": "testID", "partitionKey": "partitionKeyValue", "field1": "value1"}`)
		id = "1"
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("Post", func() {
		It("should successfully POST a resource in Azure", func() {
			testPostResource, testPostError := Post(body)
			Expect(testPostResource).To(Not(BeNil()))
			Expect(testPostError).To(BeNil())
		})
	})

	Context("Get", func() {
		It("should successfully GET a resource from Azure", func() {
			testGetResource, testGetError := Get(id)
			Expect(testGetResource).To(Not(BeNil()))
			Expect(testGetError).To(BeNil())
		})
	})

	Context("Put", func() {
		It("should successfully PUT a resource in Azure", func() {
			// TODO: [NS] Figure out how to mock Token package
			// Set mocks
			mockResource.EXPECT().BuildURI().Return("testURL", nil).Times(1)
			mockResource.EXPECT().ResourceType().Return("testResourceType", nil).Times(1)
			mockResource.EXPECT().ResourcePath().Return("testResourcePath", nil).Times(1)
			// NOTE: "dGVzdEtleQ==" -> base64("testKey")
			mockResource.EXPECT().Key().Return("dGVzdEtleQ==", nil).Times(1)

			// TODO: [NS] Validate proper values are configured in http request passed into Do
			mockHttpClient.EXPECT().Do(gomock.AssignableToTypeOf(&http.Request{})).Return(&http.Response{
				StatusCode: 200,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(`{"key": "value"}`))),
			}, nil).Times(1)

			HTTPClient = mockHttpClient
			testPutResource, testPutError := Put(mockResource, partitionKey, body)
			Expect(testPutResource).To(Not(BeNil()))
			Expect(testPutError).To(BeNil())
		})
	})

	Context("Delete", func() {
		It("should successfully DELETE a resource in Azure", func() {
			testDeleteError := Delete(id)
			Expect(testDeleteError).To(BeNil())
		})
	})
})
