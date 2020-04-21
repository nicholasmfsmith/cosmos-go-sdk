package database_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "cosmos-go-sdk/database"
)

var _ = Describe("Database", func() {

	var testDatabase *Database
	BeforeEach(func() {
		testDatabase = New("name", "key")
	})

	Context("New", func() {
		It("should successfully return a new instance of a Database Client", func() {
			Expect(testDatabase).To(BeAssignableToTypeOf(&Database{}))
		})
	})

	Context("Get", func() {
		It("should successfully fetch an Database entity from an Azure Cosmos DB account", func() {
			testEntityId := "id"
			database, testReadError := testDatabase.Get(testEntityId)
			Expect(testReadError).To(BeNil())
			Expect(database).ToNot(BeNil())
		})
	})

	Context("Delete", func() {
		It("should successfully delete current Database entity", func() {
			testDeleteError := testDatabase.Delete()
			Expect(testDeleteError).To(BeNil())
		})
	})

	Context("URI", func() {
		It("should successfully return value of URI", func() {
			testURI := testDatabase.URI()
			Expect(testURI).To(Equal(""))
		})
	})

	Context("ResourceType", func() {
		It("should successfully return value of ResourceType", func() {
			testResourceType := testDatabase.ResourceType()
			Expect(testResourceType).To(Equal(""))
		})
	})

	Context("PartitionKey", func() {
		It("should successfully return value of PartitionKey", func() {
			testPartitionKey := testDatabase.PartitionKey()
			Expect(testPartitionKey).To(Equal(""))
		})
	})
})
