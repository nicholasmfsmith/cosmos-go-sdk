package databases_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"cosmos-go-sdk/database"
	. "cosmos-go-sdk/databases"
)

var _ = Describe("Collection", func() {
	var testClient Databases
	BeforeEach(func() {
		testClient = Client("name", "key")
	})

	Context("Client", func() {
		It("should successfully return a new instance of a Databases Client", func() {
			Expect(testClient).To(BeAssignableToTypeOf(Databases{}))
		})
	})

	Context("CreateIfNotExist", func() {
		It("should successfully create an Databases entity", func() {
			testEntityId := "id"
			testEntity, testCreationError := testClient.CreateIfNotExist(testEntityId)
			Expect(testEntity).To(BeAssignableToTypeOf(database.Entity{}))
			Expect(testCreationError).To(BeNil())
		})
	})

	Context("Create", func() {
		It("should successfully create an Database entity", func() {
			testEntityId := "id"
			testEntity, testCreationError := testClient.Create(testEntityId)
			Expect(testEntity).To(BeAssignableToTypeOf(database.Entity{}))
			Expect(testCreationError).To(BeNil())
		})
	})

	Context("List", func() {
		It("should return the list of Database entity", func() {
			testEntity, testListError := testClient.List()
			Expect(testEntity).To(BeAssignableToTypeOf([]database.Entity{}))
			Expect(testListError).To(BeNil())
		})
	})
})
