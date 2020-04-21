package token_test

import (
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "cosmos-go-sdk/rest/internal/token"
)

var _ = Describe("Token", func() {
	var testToken *Token

	BeforeEach(func() {
		testToken = &Token{}
	})

	Context("Build", func() {
		It("should successfully build a token", func() {
			// NOTE: "dGVzdEtleQ==" -> base64("testKey")
			createdTestToken, err := testToken.Build(http.MethodGet, "testResourceType", "testResourcePath", "dGVzdEtleQ==")
			Expect(err).To(BeNil())
			Expect(createdTestToken).To(Not(BeEmpty()))
		})
		It("should return an error if the provided key is not base64", func() {
			createdTestToken, err := testToken.Build(http.MethodGet, "testResourceType", "testResourcePath", "testKey")
			Expect(err).To(Not(BeNil()))
			Expect(err.Error()).To(ContainSubstring("Error decoding provided key"))
			Expect(createdTestToken).To(BeEmpty())
		})
		It("should return a token that is QueryEscaped", func() {
			// NOTE: "dGVzdEtleQ==" -> base64("testKey")
			createdTestToken, err := testToken.Build(http.MethodGet, "testResourceType", "testResourcePath", "dGVzdEtleQ==")
			Expect(err).To(BeNil())
			Expect(createdTestToken).To(ContainSubstring("type%3Dmaster%26ver%3D1.0%26sig%3D"))
		})
	})
})
