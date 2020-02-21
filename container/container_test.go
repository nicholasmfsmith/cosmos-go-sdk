package container_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "cosmos-go-sdk/container"
)

var _ = Describe("Container", func() {

	Context("Client", func() {
		It("should successfully return a new instance of a Container Client", func() {
			testClient := Client("name", "dbName", "key")
			Expect(testClient).To(BeAssignableToTypeOf(Container{}))
		})
	})

	Context("Get", func() {
		It("should successfully fetch an Container Document", func() {
			testClient := Client("name", "dbName", "key")
			container, testReadError := testClient.Get()
			Expect(testReadError).To(BeNil())
			Expect(container).ToNot(BeNil())
		})
	})

	Context("Delete", func() {
		It("should successfully fetch an Container Document", func() {
			testClient := Client("name", "dbName", "key")
			testDeleteError := testClient.Delete()
			Expect(testDeleteError).To(BeNil())
		})
	})

	Context("Replace", func() {
		It("should successfully fetch an Container Document", func() {
			testClient := Client("name", "dbName", "key")
			testEntity := Entity{}
			container, testReplaceError := testClient.Replace(testEntity)
			Expect(testReplaceError).To(BeNil())
			Expect(container).ToNot(BeNil())
		})
	})
})
