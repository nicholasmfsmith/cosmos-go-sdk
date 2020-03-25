package token_test

import (
	"net/http"
	"strings"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "cosmos-go-sdk/token"
)

var _ = Describe("Token", func() {
	var testToken *Token

	BeforeEach(func() {
		// NOTE: "dGVzdEtleQ==" -> base64("testKey")
		testToken = New(http.MethodGet, "testResourceType", "testResourceID", "dGVzdEtleQ==")
	})

	Context("New", func() {
		It("should return a valid Token pointer", func() {
			Expect(testToken).To(BeAssignableToTypeOf(&Token{}))
		})
		It("should return a valid Token pointer with a missing token", func() {
			testToken := New(http.MethodGet, "testResourceType", "testResourceID", "testKey")
			Expect(testToken.Method).To(Not(BeEmpty()))
			Expect(testToken.ResourceType).To(Not(BeEmpty()))
			Expect(testToken.ResourceID).To(Not(BeEmpty()))
			Expect(testToken.Key).To(Not(BeEmpty()))
			Expect(testToken.Date).To(Not(BeEmpty()))
			Expect(testToken.Token).To(BeEmpty())
		})
		It("should return a valid Token pointer with a date that satisfies the correct standard", func() {
			_, err := time.Parse(http.TimeFormat, strings.ToUpper(testToken.Date))
			Expect(err).To(BeNil())
		})
	})

	Context("Build", func() {
		It("should successfully build a token", func() {
			err := testToken.Build()
			Expect(err).To(BeNil())
		})
		It("should return an error if the provided key is not base64", func() {
			testToken.Key = "testKey"
			err := testToken.Build()
			Expect(err).To(Not(BeNil()))
			Expect(err.Error()).To(ContainSubstring("Error decoding provided key"))
		})
		It("should return an error if the provided HTTP Method is empty", func() {
			testToken.Method = ""
			err := testToken.Build()
			Expect(err).To(Not(BeNil()))
			Expect(err.Error()).To(Equal("Error building token: Provided HTTP method is empty"))
		})
		It("should return an error if the provided Resource Type is empty", func() {
			testToken.ResourceType = ""
			err := testToken.Build()
			Expect(err).To(Not(BeNil()))
			Expect(err.Error()).To(Equal("Error building token: Provided Resource Type is empty"))
		})
		It("should return an error if the provided Resource ID is empty", func() {
			testToken.ResourceID = ""
			err := testToken.Build()
			Expect(err).To(Not(BeNil()))
			Expect(err.Error()).To(Equal("Error building token: Provided Resource ID is empty"))
		})
		It("should return an error if the provided Key is empty", func() {
			testToken.Key = ""
			err := testToken.Build()
			Expect(err).To(Not(BeNil()))
			Expect(err.Error()).To(Equal("Error building token: Provided key is empty"))
		})
		It("should return an error if the provided Date is empty", func() {
			testToken.Date = ""
			err := testToken.Build()
			Expect(err).To(Not(BeNil()))
			Expect(err.Error()).To(Equal("Error building token: Provided date is empty"))
		})
		It("should return a token that is QueryEscaped", func() {
			err := testToken.Build()
			Expect(err).To(BeNil())
			Expect(testToken.Token).To(ContainSubstring("type%3Dmaster%26ver%3D1.0%26sig%3D"))
		})
	})
})
