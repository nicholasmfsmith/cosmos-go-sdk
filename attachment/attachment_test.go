package attachment_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "cosmos-go-sdk/attachment"
)

var _ = Describe("Attachment", func() {
	var testAttachment *Attachment
	var partitionKey string

	BeforeEach(func() {
		partitionKey = "key"
		testAttachment = New(partitionKey)
	})

	Context("New", func() {
		It("return a new instance of an Attachment", func() {
			Expect(testAttachment).To(BeAssignableToTypeOf(&Attachment{}))
		})
	})

	Context("URI", func() {
		It("should successfully return value of URI", func() {
			testURI := testAttachment.URI()
			Expect(testURI).To(Equal(""))
		})
	})

	Context("ResourceType", func() {
		It("should successfully return value of URI", func() {
			testResourceType := testAttachment.ResourceType()
			Expect(testResourceType).To(Equal(""))
		})
	})

	Context("PartitionKey", func() {
		It("should successfully return value of URI", func() {
			testPartitionKey := testAttachment.PartitionKey()
			Expect(testPartitionKey).To(Equal(""))
		})
	})
})
