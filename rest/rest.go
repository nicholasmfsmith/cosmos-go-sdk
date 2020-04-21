// Package rest provides functionality to handle
// the HTTP interactions with the Azure API.
// TODO: Implementation of HTTP specific errors
package rest

import (
	"bytes"
	"cosmos-go-sdk/rest/internal/token"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// TODO: [NS] Make Token more mockable

const (
	// Current version of the Azure Cosmos API used
	apiVersion = "2017-02-22"
	// TODO: [NS] Find better ways to handle timeouts
	// Should this be an input parameters?
	// Different per HTTP request type?
	requestTimeout = 30 * time.Second
)

var (
	// HTTPClient is the shared HTTP client used for all requests
	HTTPClient IHttpClient
)

// IHttpClient is an interface used to override the real HTTP client for testing
type IHttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// IResource is an interface to define the contract between any resource and this package.
// All resources must be able to build their URI and provide their resource type, path, and primary key
type IResource interface {
	URI() string
	ResourceType() string
	PartitionKey() string
}

// TODO: [NS] Explore better ways to override HTTP client for testing
func init() {
	// Initialize the HttpClient to the true http.Client type, unless overridden by a mock for testing
	HTTPClient = &http.Client{
		Timeout: requestTimeout,
	}
}

// Post performs a POST HTTP request to the Azure API to create the provided resource.
// It returns the created resource as a byte array and any errors encountered.
func Post(resource []byte) ([]byte, error) {
	return []byte(""), nil
}

// Get performs a GET HTTP request to the Azure API to read the resource
// identified by the provided resource ID.
// It returns the requested resource as a byte array and any errors encountered.
func Get(resource IResource, key string) ([]byte, error) {

	uri := resource.URI()
	resourceType := resource.ResourceType()
	resourcePath := extractResourcePathFromURI(uri)
	partitionKey := resource.PartitionKey()

	// Get token, if any error, return immediately
	requestToken := &token.Token{}
	currentToken, requestTokenBuildErr := requestToken.Build(http.MethodGet, resourceType, resourcePath, key)
	if requestTokenBuildErr != nil {
		return nil, requestTokenBuildErr
	}

	request, errNewRequest := http.NewRequest(http.MethodGet, uri, nil)
	if errNewRequest != nil {
		return nil, errNewRequest
	}

	// TODO: Adding optional headers in a separate PR
	request.Header["authorization"] = []string{currentToken}
	request.Header["x-ms-documentdb-partitionkey"] = []string{partitionKey}
	request.Header["x-ms-version"] = []string{apiVersion}
	request.Header["x-ms-date"] = []string{strings.ToLower(time.Now().UTC().Format(http.TimeFormat))}

	response, errRequest := HTTPClient.Do(request)
	if errRequest != nil {
		return nil, errRequest
	}
	defer response.Body.Close()

	responseBody, errReadResponseBody := ioutil.ReadAll(response.Body)
	if errReadResponseBody != nil {
		return nil, errReadResponseBody
	}

	return responseBody, nil
}

// Put performs a PUT HTTP request to the Azure API for the provided
// resource to replace the existing remote resource with the provided resource.
// It returns the updated resource as a byte array and any errors encountered.
// TODO: [NS] How should partitionKey be handled? Should it be optional?
// TODO: [NS] Add better error messages
func Put(resource IResource, key string, body []byte) ([]byte, error) {
	// Get required data points from calling resource, if error return immediately
	uri := resource.URI()
	resourceType := resource.ResourceType()
	resourcePath := extractResourcePathFromURI(uri)

	// Get token, if any error, return immediately
	requestToken := &token.Token{}
	currentToken, requestTokenBuildErr := requestToken.Build(http.MethodPut, resourceType, resourcePath, key)
	if requestTokenBuildErr != nil {
		return nil, requestTokenBuildErr
	}

	// Notice the format required for the partition Key
	partitionKey := fmt.Sprintf(`["%s"]`, resource.PartitionKey())

	// Create request
	req, newRequestErr := http.NewRequest(http.MethodPut, uri, bytes.NewBuffer(body))
	if newRequestErr != nil {
		return nil, newRequestErr
	}

	// Assign required headers
	req.Header["x-ms-documentdb-partitionkey"] = []string{partitionKey}
	req.Header["x-ms-version"] = []string{apiVersion}
	req.Header["x-ms-date"] = []string{strings.ToLower(time.Now().UTC().Format(http.TimeFormat))}
	req.Header["authorization"] = []string{currentToken}
	req.Header["content-type"] = []string{"application/json"}

	// TODO: [NS] Handle optional headers

	resp, requestErr := HTTPClient.Do(req)
	if requestErr != nil {
		return nil, requestErr
	}
	defer resp.Body.Close()
	respBody, readRespBodyErr := ioutil.ReadAll(resp.Body)
	if readRespBodyErr != nil {
		return nil, readRespBodyErr
	}
	return respBody, nil
}

// Delete performs a DELETE HTTP request to the Azure API to remove the resource
// identified by the provided resource ID.
// It returns any errors encountered.
func Delete(id string) error {
	return nil
}

// Note: URI follows the below format:
// https://{databaseaccount}.documents.azure.com/dbs/{db-id}/colls/{coll-id}/docs/{doc-name}
// TODO: [NS] Add unit tests
func extractResourcePathFromURI(uri string) string {
	res := strings.Split(uri, ".com")
	if len(res) < 2 {
		return ""
	}
	return res[1]
}
