package client_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "cosmos-go-sdk/client"
)

var _ = Describe("Client", func() {
	Context("CosmosClient", func() {
		It("should successfully return a new instance of Client", func() {
			testClient := CosmosClient("www.testurl.com", "aTESTKEy")
			Expect(testClient).To(BeAssignableToTypeOf(Client{}))
		})
	})
})
