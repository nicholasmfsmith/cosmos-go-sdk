package database_test

import (
	. "cosmos-go-sdk/database"
	"cosmos-go-sdk/testdata/mocks"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Database", func() {
	var (
		mockCtrl     *gomock.Controller
		mockRequest  *mocks.MockIRequest
		testDatabase Database
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockRequest = mocks.NewMockIRequest(mockCtrl)
		testDatabase = Database{
			URI:     "testID",
			Name:    "testPartitionKey",
			Key:     "testKey",
			Request: mockRequest,
		}
	})

	Context("New", func() {
		It("should successfully return a new instance of a Database Client", func() {
			testDatabase = New("testDb", "KEY", "localhost")
			Expect(testDatabase).To(BeAssignableToTypeOf(Database{}))
		})

		It("should successfully return a new instance of a Database with a type Request as a property", func() {
			testDatabase = New("testDb", "KEY", "localhost")
		})

		It("should successfully return a new instance of a Database with modified uri", func() {
			baseUri := "localhost"
			databaseName := "testDb"
			path := "/dbs/" + databaseName
			testDatabase = New(databaseName, "KEY", baseUri)
			Expect(testDatabase.URI).To(Equal(baseUri + path))
		})

		It("should successfully return a new instance of a Database with passed key", func() {
			key := "KEY"
			testDatabase = New("testDb", key, "localhost")
			Expect(testDatabase.Key).To(Equal(key))
		})

		It("should successfully return a new instance of a Database with passed name", func() {
			name := "testDb"
			testDatabase = New(name, "KEY", "localhost")
			Expect(testDatabase.Name).To(Equal(name))
		})
	})

	Context("Read", func() {
		It("should successfully fetch an Database entity", func() {
			azureRecord := struct {
				ID    string
				RID   string
				TS    int64
				SELF  string
				ETAG  string
				COLLS string
				USER  string
			}{
				"ID",
				"RID",
				1000,
				"SELF",
				"ETAG",
				"COLLS",
				"USER",
			}
			azureRecordString := fmt.Sprintf(`{"id": "%s", "_rid": "%s", "_ts": %d, "_self": "%s", "_etag": "%s", "_colls": "%s", "_user": "%s"}`, azureRecord.ID, azureRecord.RID, azureRecord.TS, azureRecord.SELF, azureRecord.ETAG, azureRecord.COLLS, azureRecord.USER)
			mockRequest.EXPECT().Get().Return([]byte(azureRecordString), nil).Times(1)
			database, testReadError := testDatabase.Read()
			Expect(testReadError).ShouldNot(HaveOccurred())
			Expect(database).To(Equal(Entity(azureRecord)))
		})

		It("should return error if unable to successfully fetch an database entity", func() {
			mockRequest.EXPECT().Get().Return(nil, errors.New("http error")).Times(1)
			database, testReadError := testDatabase.Read()
			Expect(database).To(Equal(Entity{
				ID:    "",
				RID:   "",
				TS:    0,
				SELF:  "",
				ETAG:  "",
				COLLS: "",
				USER:  "",
			}))
			Expect(testReadError).Should(HaveOccurred())
		})
	})

	Context("Delete", func() {
		It("should successfully delete current Database entity", func() {
			testDeleteError := testDatabase.Delete()
			Expect(testDeleteError).To(BeNil())
		})
	})
})
