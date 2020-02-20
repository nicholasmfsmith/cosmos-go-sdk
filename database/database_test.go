package database_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "cosmos-go-sdk/database"
)

var _ = Describe("Database", func() {

	Context("Create", func() {
		It("should successfully create an database item on an Azure Cosmos DB account", func() {
			testDocumentId := "id"
			testClient := Client("testDB", "testKey")
			testCreateError := testClient.Create(testDocumentId)
			Expect(testCreateError).To(BeNil())
		})
	})

	Context("Get", func() {
		It("should successfully fetch an database item from an Azure Cosmos DB account", func() {
			testDocumentId := "id"
			testClient := Client("testDB", "testKey")
			database, testReadError := testClient.Get(testDocumentId)
			Expect(testReadError).To(BeNil())
			Expect(database).ToNot(BeNil())
		})
	})

	Context("List", func() {
		It("should successfully fetch an database item from an Azure Cosmos DB account", func() {
			testClient := Client("testDB", "testKey")
			databases, testReadError := testClient.List()
			Expect(testReadError).To(BeNil())
			Expect(databases).ToNot(BeNil())
		})
	})

	Context("Delete", func() {
		It("should successfully delete an item from the Azure Cosmos Database Container", func() {
			testDocumentId := "id"
			testClient := Client("testDB", "testKey")
			testDeleteError, testReadError := testClient.Delete(testDocumentId)
			Expect(testReadError).To(BeNil())
			Expect(testDeleteError).To(BeNil())
		})
	})
})
