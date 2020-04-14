// Package rest provides functionality to handle
// the HTTP interactions with the Azure API.
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

// TODO: [NS] Explore better ways to override HTTP client for testing

// IHttpClient is an interface used to override the real HTTP client for testing
type IHttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// IResource is an interface to define the contract between any resource and this package.
// All resources must be able to build their URI and provide their resource type, path, and primary key
type IResource interface {
	BuildURI() (string, error)
	ResourceType() (string, error)
	ResourcePath() (string, error)
	Key() (string, error)
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
func Get(id string) ([]byte, error) {
	return []byte(""), nil
}

// Put performs a PUT HTTP request to the Azure API for the provided
// resource to replace the existing remote resource with the provided resource.
// It returns the updated resource as a byte array and any errors encountered.
// TODO: [NS] How should partitionKey be handled? Should it be optional?
// TODO: [NS] Add better error messages
func Put(resource IResource, partitionKey string, body []byte) ([]byte, error) {
	// Get required data points from calling resource, if error return immediately
	uri, buildURIErr := resource.BuildURI()
	if buildURIErr != nil {
		return nil, buildURIErr
	}
	resourceType, resourceTypeErr := resource.ResourceType()
	if resourceTypeErr != nil {
		return nil, resourceTypeErr
	}
	resourcePath, resourcePathErr := resource.ResourcePath()
	if resourcePathErr != nil {
		return nil, resourcePathErr
	}
	key, keyErr := resource.Key()
	if keyErr != nil {
		return nil, keyErr
	}

	// Get token, if any error, return immediately
	requestToken := token.New(http.MethodPut, resourceType, resourcePath, key)
	requestTokenBuildErr := requestToken.Build()
	if requestTokenBuildErr != nil {
		return nil, requestTokenBuildErr
	}

	// Notice the format required for the partition Key
	partitionKey = fmt.Sprintf(`["%s"]`, partitionKey)

	// Create request
	req, newRequestErr := http.NewRequest(http.MethodPut, uri, bytes.NewBuffer(body))
	if newRequestErr != nil {
		return nil, newRequestErr
	}

	// Assign required headers
	req.Header["x-ms-documentdb-partitionkey"] = []string{partitionKey}
	req.Header["x-ms-version"] = []string{apiVersion}
	req.Header["x-ms-date"] = []string{strings.ToLower(time.Now().UTC().Format(http.TimeFormat))}
	req.Header["authorization"] = []string{requestToken.Token}
	req.Header["content-type"] = []string{"application/json"}

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
