package database_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "cosmos-go-sdk/database"
)

var _ = Describe("Database", func() {

	var testClient Database
	BeforeEach(func() {
		testClient = Client("name", "key")
	})

	Context("Client", func() {
		It("should successfully return a new instance of a Database Client", func() {
			Expect(testClient).To(BeAssignableToTypeOf(Database{}))
		})
	})

	Context("Get", func() {
		It("should successfully fetch an Database entity from an Azure Cosmos DB account", func() {
			testEntityId := "id"
			database, testReadError := testClient.Get(testEntityId)
			Expect(testReadError).To(BeNil())
			Expect(database).ToNot(BeNil())
		})
	})

	Context("Delete", func() {
		It("should successfully delete current Database entity", func() {
			testDeleteError := testClient.Delete()
			Expect(testDeleteError).To(BeNil())
		})
	})
})
