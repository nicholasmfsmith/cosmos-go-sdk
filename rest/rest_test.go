package rest_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "cosmos-go-sdk/rest"
)

var _ = Describe("Rest", func() {
	var resource []byte
	var id string

	BeforeEach(func() {
		resource = []byte("This is a test resource")
		id = "1"
	})

	Context("Post", func() {
		It("should successfully POST a resource in Azure", func() {
			testPostResource, testPostError := Post(resource)
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
			testPutResource, testPutError := Put(id, resource)
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
