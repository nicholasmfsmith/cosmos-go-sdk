// Package rest provides functionality to handle
// the HTTP interactions with the Azure API.
// TODO: Implementation of HTTP specific errors
package rest

import (
	"bytes"
	"cosmos-go-sdk/rest/internal/token"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
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

// AzureHTTPError represents the http error schema that Azure uses for it's api
type AzureHTTPError struct {
	Code    string
	Message string
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
	resourcePath, invalidUriError := extractResourcePathFromURI(request.URI)
	if invalidUriError != nil {
		return nil, invalidUriError
	}

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

	azureHTTPError, unmarshalError := azureHTTPErrorCheck(response.StatusCode, []byte(responseBody))
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	if azureHTTPError.Code != "" {
		return nil, fmt.Errorf("code: %s, message: %s", azureHTTPError.Code, azureHTTPError.Message)
	}

	return responseBody, nil
}

func azureHTTPErrorCheck(statusCode int, responseBody []byte) (AzureHTTPError, error) {
	var azureHTTPError AzureHTTPError
	if statusCode >= 400 {
		unmarshalError := json.Unmarshal(responseBody, &azureHTTPError)
		if unmarshalError != nil {
			return azureHTTPError, fmt.Errorf("unknown error schema : %s", string(responseBody))
		}
	}
	return azureHTTPError, nil
}

// Put performs a PUT HTTP request to the Azure API for the provided
// resource to replace the existing remote resource with the provided resource.
// It returns the updated resource as a byte array and any errors encountered.
// TODO: [NS] How should partitionKey be handled? Should it be optional?
// TODO: [NS] Add better error messages
func (request Request) Put(resource []byte) ([]byte, error) {
	resourcePath, invalidUriError := extractResourcePathFromURI(request.URI)
	if invalidUriError != nil {
		return nil, invalidUriError
	}

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

	response, requestErr := request.HTTP.Do(currentHTTPRequest)
	if requestErr != nil {
		return nil, requestErr
	}
	defer response.Body.Close()
	responseBody, readRespBodyErr := ioutil.ReadAll(response.Body)

	azureHTTPError, unmarshalError := azureHTTPErrorCheck(response.StatusCode, []byte(responseBody))
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	if azureHTTPError.Code != "" {
		return nil, fmt.Errorf("code: %s, message: %s", azureHTTPError.Code, azureHTTPError.Message)
	}

	if readRespBodyErr != nil {
		return nil, readRespBodyErr
	}
	return responseBody, nil
}

// Delete performs a DELETE HTTP request to the Azure API to remove the resource
// identified by the provided resource ID.
// It returns any errors encountered.
func (request Request) Delete() error {
	return nil
}

// Note: URI follows the below format:
// https://{databaseaccount}.documents.azure.com/dbs/{db-id}/colls/{coll-id}/docs/{doc-name}
// TODO: [SC] Add unit tests
func extractResourcePathFromURI(uri string) (string, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return "", err
	}
	path := u.Path
	// Note: We could have the following 3 cases
	// {host}/{path} in this case we return path
	// {host}/ in this case we return the empty string
	// {host} in this case we return the empty string
	if path[0:1] == "/" {
		path = path[1:]
	}

	return path, nil
}
