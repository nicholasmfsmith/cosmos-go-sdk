// Package rest provides functionality to handle
// the HTTP interactions with the Azure API.
// TODO: Implementation of HTTP specific errors
package rest

import (
	"bytes"
	"cosmos-go-sdk/rest/internal/token"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	// Current version of the Azure Cosmos API used
	apiVersion = "2017-02-22"
	// TODO: [NS] Find better ways to handle timeouts
	requestTimeout = 30 * time.Second
)

// IRequest is the interface that defines the functionality of the rest package
type IRequest interface {
	Get() ([]byte, error)
	Post(resource []byte) ([]byte, error)
	Put(resource []byte) ([]byte, error)
	Delete() error
}

// Request is an implementation of the IRequest interface for the Azure API
type Request struct {
	URI          string
	ResourceType string
	Key          string
	HTTP         IHttpClient
	Token        token.IToken
}

// IHttpClient is an interface used to override the real HTTP client for testing
// TODO: [NS] This is only here because I couldn't find an interface in the http package to leverage.
// TODO: [NS] Do more investigation to see if there is an existing interface we can use
type IHttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// New creates a fresh instance of a request based on the provided parameters
// with the necessary HTTP and token clients
func New(uri, resourceType, key string) Request {
	return Request{
		uri,
		resourceType,
		key,
		&http.Client{
			Timeout: requestTimeout,
		},
		&token.Token{},
	}
}

// Post performs a POST HTTP request to the Azure API to create the provided resource.
// It returns the created resource as a byte array and any errors encountered.
func (request Request) Post(resource []byte) ([]byte, error) {
	return []byte(""), nil
}

// Get performs a GET HTTP request to the Azure API to read the resource
// identified by the provided resource ID.
// It returns the requested resource as a byte array and any errors encountered.
func (request Request) Get() ([]byte, error) {
	resourcePath := extractResourcePathFromURI(request.URI)

	// Get token, if any error, return immediately
	currentToken, requestTokenBuildErr := request.Token.Build(http.MethodGet, request.ResourceType, resourcePath, request.Key)
	if requestTokenBuildErr != nil {
		return nil, requestTokenBuildErr
	}

	currentHTTPRequest, errNewRequest := http.NewRequest(http.MethodGet, request.URI, nil)
	if errNewRequest != nil {
		return nil, errNewRequest
	}

	// TODO: Adding optional headers in a separate PR
	currentHTTPRequest.Header["authorization"] = []string{currentToken}
	// TODO: [NS] Figure out how to handle partition key
	// currentHTTPRequest.Header["x-ms-documentdb-partitionkey"] = []string{partitionKey}
	currentHTTPRequest.Header["x-ms-version"] = []string{apiVersion}
	currentHTTPRequest.Header["x-ms-date"] = []string{strings.ToLower(time.Now().UTC().Format(http.TimeFormat))}

	response, errRequest := request.HTTP.Do(currentHTTPRequest)
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
func (request Request) Put(resource []byte) ([]byte, error) {
	resourcePath := extractResourcePathFromURI(request.URI)

	// Get token, if any error, return immediately
	currentToken, requestTokenBuildErr := request.Token.Build(http.MethodPut, request.ResourceType, resourcePath, request.Key)
	if requestTokenBuildErr != nil {
		return nil, requestTokenBuildErr
	}

	// TODO: [NS] Figure out how to handle partition key
	// Notice the format required for the partition Key
	// partitionKey := fmt.Sprintf(`["%s"]`, resource.PartitionKey())

	// Create request
	currentHTTPRequest, newRequestErr := http.NewRequest(http.MethodPut, request.URI, bytes.NewBuffer(resource))
	if newRequestErr != nil {
		return nil, newRequestErr
	}

	// Assign required headers
	// currentHTTPRequest.Header["x-ms-documentdb-partitionkey"] = []string{partitionKey}
	currentHTTPRequest.Header["x-ms-version"] = []string{apiVersion}
	currentHTTPRequest.Header["x-ms-date"] = []string{strings.ToLower(time.Now().UTC().Format(http.TimeFormat))}
	currentHTTPRequest.Header["authorization"] = []string{currentToken}
	currentHTTPRequest.Header["content-type"] = []string{"application/json"}

	// TODO: [NS] Handle optional headers

	resp, requestErr := request.HTTP.Do(currentHTTPRequest)
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
func (request Request) Delete() error {
	return nil
}

// Note: URI follows the below format:
// https://{databaseaccount}.documents.azure.com/dbs/{db-id}/colls/{coll-id}/docs/{doc-name}
// TODO: [NS] Add unit tests
func extractResourcePathFromURI(uri string) string {
	res := strings.Split(uri, ".com/")
	if len(res) < 2 {
		return ""
	}
	return res[1]
}
