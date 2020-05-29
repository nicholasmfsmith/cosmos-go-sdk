package rest_test

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "cosmos-go-sdk/rest"
	"cosmos-go-sdk/testdata/mocks"
)

var _ = Describe("Rest", func() {
	var (
		mockCtrl         *gomock.Controller
		mockHttpClient   *mocks.MockIHttpClient
		mockToken        *mocks.MockIToken
		resource         []byte
		testRequest      Request
		testResourceLink string
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockHttpClient = mocks.NewMockIHttpClient(mockCtrl)
		mockToken = mocks.NewMockIToken(mockCtrl)
		resource = []byte(`{"id": "testID", "partitionKey": "partitionKeyValue", "field1": "value1"}`)
		testRequest = Request{
			URI:          "https://mock-test-database-account.documents.azure.com/dbs/{db-id}/colls/{coll-id}/docs/{doc-name}",
			ResourceType: "test",
			Key:          "testKey",
			HTTP:         mockHttpClient,
			Token:        mockToken,
		}
		testResourceLink = "dbs/{db-id}/colls/{coll-id}/docs/{doc-name}"
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("New", func() {
		It("should successfully create a non-null, fresh instance of Request", func() {
			testRequest = New("testURI", "testResourceType", "testKey")
			Expect(testRequest).To(Not(BeNil()))
			Expect(testRequest).To(BeAssignableToTypeOf(Request{}))
		})
		It("should successfully create a fresh instance of Request with the correct URI", func() {
			testRequest = New("testURI", "testResourceType", "testKey")
			Expect(testRequest.URI).To(Equal("testURI"))
		})
		It("should successfully create a fresh instance of Request with the correct ResourceType", func() {
			testRequest = New("testURI", "testResourceType", "testKey")
			Expect(testRequest.ResourceType).To(Equal("testResourceType"))
		})
		It("should successfully create a fresh instance of Request with the correct Key", func() {
			testRequest = New("testURI", "testResourceType", "testKey")
			Expect(testRequest.Key).To(Equal("testKey"))
		})
	})

	Context("Post", func() {
		It("should successfully POST a resource in Azure", func() {
			// TODO: [NS] Validate proper values are configured in http request passed into Do
			mockHttpClient.EXPECT().Do(gomock.AssignableToTypeOf(&http.Request{})).Return(&http.Response{
				StatusCode: http.StatusCreated,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(`{"key": "value"}`))),
			}, nil).Times(1)
			mockToken.EXPECT().Build(http.MethodPost, testRequest.ResourceType, testResourceLink, testRequest.Key).Return("testToken", nil).Times(1)

			testPostResource, testPostError := testRequest.Post(resource)
			Expect(testPostResource).To(Equal([]byte(`{"key": "value"}`)))
			Expect(testPostError).To(BeNil())
		})

		It("should successfully handle errors when performing a POST for a resource in Azure", func() {
			testError := errors.New("This is a test HTTP client error")
			mockHttpClient.EXPECT().Do(gomock.AssignableToTypeOf(&http.Request{})).Return(nil, testError).Times(1)
			mockToken.EXPECT().Build(http.MethodPost, testRequest.ResourceType, testResourceLink, testRequest.Key).Return("testToken", nil).Times(1)

			testPostResource, testPostError := testRequest.Post(resource)
			Expect(testPostResource).To(BeNil())
			Expect(testPostError.Error()).To(Equal("This is a test HTTP client error"))
		})
	})

	Context("Get", func() {
		It("should successfully GET a resource from Azure", func() {
			mockHttpClient.EXPECT().Do(gomock.AssignableToTypeOf(&http.Request{})).Return(&http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(`{"key": "value"}`))),
			}, nil).Times(1)
			mockToken.EXPECT().Build(http.MethodGet, testRequest.ResourceType, testResourceLink, testRequest.Key).Return("testToken", nil).Times(1)

			testGetResource, testGetError := testRequest.Get()
			Expect(testGetResource).To(Equal([]byte(`{"key": "value"}`)))
			Expect(testGetError).To(BeNil())
		})
	})

	Context("Put", func() {
		// TODO: [NS] ADD MORE TESTS!
		It("should successfully PUT a resource in Azure", func() {
			// TODO: [NS] Validate proper values are configured in http request passed into Do
			mockHttpClient.EXPECT().Do(gomock.AssignableToTypeOf(&http.Request{})).Return(&http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(`{"key": "value"}`))),
			}, nil).Times(1)
			mockToken.EXPECT().Build(http.MethodPut, testRequest.ResourceType, testResourceLink, testRequest.Key).Return("testToken", nil).Times(1)
			testPutResource, testPutError := testRequest.Put(resource)
			Expect(testPutResource).To(Equal([]byte(`{"key": "value"}`)))
			Expect(testPutError).To(BeNil())
		})
	})

	Context("Delete", func() {
		It("should successfully DELETE a resource in Azure", func() {
			mockHttpClient.EXPECT().Do(gomock.AssignableToTypeOf(&http.Request{})).Return(&http.Response{
				StatusCode: http.StatusNoContent,
			}, nil).Times(1)
			mockToken.EXPECT().Build(http.MethodDelete, testRequest.ResourceType, testResourceLink, testRequest.Key).Return("testToken", nil).Times(1)
			testDeleteError := testRequest.Delete()
			Expect(testDeleteError).To(BeNil())
		})
	})
})
