package client_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "cosmos-go-sdk/client"
)

var _ = Describe("Client", func() {
	Context("New", func() {
		It("should successfully return a new instance of Client", func() {
			testClient := New("www.testurl.com", "aTESTKEy")
			Expect(testClient).To(BeAssignableToTypeOf(Client{}))
		})
	})
})
