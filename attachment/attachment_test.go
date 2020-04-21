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

	Context("URI", func() {
		It("should successfully return value of URI", func() {
			Expect(testAttachment).To(BeAssignableToTypeOf(&Attachment{}))
			testURI := testAttachment.URI()
			Expect(testURI).To(Equal(""))
		})
	})

	Context("ResourceType", func() {
		It("should successfully return value of URI", func() {
			Expect(testAttachment).To(BeAssignableToTypeOf(&Attachment{}))
			testResourceType := testAttachment.ResourceType()
			Expect(testResourceType).To(Equal(""))
		})
	})

	Context("PartitionKey", func() {
		It("should successfully return value of URI", func() {
			Expect(testAttachment).To(BeAssignableToTypeOf(&Attachment{}))
			testPartitionKey := testAttachment.PartitionKey()
			Expect(testPartitionKey).To(Equal(""))
		})
	})
})
