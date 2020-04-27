package client_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "cosmos-go-sdk/client"
	"cosmos-go-sdk/database"
)

var _ = Describe("Client", func() {

	var testClient *Client
	BeforeEach(func() {
		testClient = New("url", "key")
	})

	Context("New", func() {
		It("should successfully return a new instance of Client", func() {
			Expect(testClient).To(BeAssignableToTypeOf(&Client{}))
		})
	})

	Context("Database", func() {
		It("should successfully return a new instance of Database with the current instance of Client", func() {
			db := testClient.Database("name")
			Expect(db).To(BeAssignableToTypeOf(database.Database{}))
		})
	})
})
