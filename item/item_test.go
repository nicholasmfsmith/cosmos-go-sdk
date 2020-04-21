package item_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "cosmos-go-sdk/item"
)

var _ = Describe("Item", func() {
	var testItem Item

	BeforeEach(func() {
		testItem = Item{}
	})

	Context("New", func() {
		It("should successfully return a new instance of Item", func() {
			testItem = New("testId", "testPartitionKey", "testDatabaseName", "testContainerName", "testKey")
			Expect(testItem).To(BeAssignableToTypeOf(Item{}))
		})
	})

	Context("Create", func() {
		It("should successfully create an item in the Azure Cosmos Database Container", func() {
			testDocument := []byte("This is a test new document")
			testCreateError := testItem.Create(testDocument)
			Expect(testCreateError).To(BeNil())
		})
	})

	// TODO: [NS] Fully implement item.Read(). Doing so will make this test pass
	Context("Read", func() {
		It("should successfully read an item from the Azure Cosmos Database Container", func() {
			document, testReadError := testItem.Read()
			Expect(testReadError).To(BeNil())
			Expect(document).ToNot(BeNil())
		})
	})

	Context("Update", func() {
		It("should successfully update an item in the Azure Cosmos Database Container", func() {
			testDocument := []byte("This is a test updated document")
			testUpdateError := testItem.Update(testDocument)
			Expect(testUpdateError).To(BeNil())
		})
	})

	Context("Delete", func() {
		It("should successfully delete an item from the Azure Cosmos Database Container", func() {
			testDeleteError := testItem.Delete()
			Expect(testDeleteError).To(BeNil())
		})
	})

	Context("URI", func() {
		It("should successfully return value of URI", func() {
			testURI := testItem.URI()
			Expect(testURI).To(Equal(""))
		})
	})

	Context("ResourceType", func() {
		It("should successfully return value of ResourceType", func() {
			testResourceType := testItem.ResourceType()
			Expect(testResourceType).To(Equal(""))
		})
	})

	Context("PartitionKey", func() {
		It("should successfully return value of PartitionKey", func() {
			testPartitionKey := testItem.PartitionKey()
			Expect(testPartitionKey).To(Equal(""))
		})
	})
})
