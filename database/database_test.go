package database_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "cosmos-go-sdk/database"
)

var _ = Describe("Database", func() {

	Context("Client", func() {
		It("should successfully return a new instance of a Database Client", func() {
			testClient := Client("name", "key")
			Expect(testClient).To(BeAssignableToTypeOf(Database{}))
		})
	})

	Context("Get", func() {
		It("should successfully fetch an Database entity from an Azure Cosmos DB account", func() {
			testEntityId := "id"
			testClient := Client("testDB", "testKey")
			database, testReadError := testClient.Get(testEntityId)
			Expect(testReadError).To(BeNil())
			Expect(database).ToNot(BeNil())
		})
	})

	Context("Delete", func() {
		It("should successfully delete current Database entity", func() {
			testClient := Client("testDB", "testKey")
			testDeleteError, testReadError := testClient.Delete()
			Expect(testReadError).To(BeNil())
			Expect(testDeleteError).To(BeNil())
		})
	})
})
