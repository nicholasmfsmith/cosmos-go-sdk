package offer_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "cosmos-go-sdk/offer"
)

var _ = Describe("Offer", func() {
	var testOffer *Offer
	var partitionKey string

	BeforeEach(func() {
		partitionKey = "key"
		testOffer = New(partitionKey)
	})

	Context("New", func() {
		It("return a new instance of an Offer", func() {
			Expect(testOffer).To(BeAssignableToTypeOf(&Offer{}))
		})
	})

	Context("URI", func() {
		It("should successfully return value of URI", func() {
			testURI := testOffer.URI()
			Expect(testURI).To(Equal(""))
		})
	})

	Context("ResourceType", func() {
		It("should successfully return value of ResourceType", func() {
			testResourceType := testOffer.ResourceType()
			Expect(testResourceType).To(Equal(""))
		})
	})

	Context("PartitionKey", func() {
		It("should successfully return value of PartitionKey", func() {
			testPartitionKey := testOffer.PartitionKey()
			Expect(testPartitionKey).To(Equal(""))
		})
	})
})
