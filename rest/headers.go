/*
Package rest headers source file addresses the request headers used per spec  https://docs.microsoft.com/en-us/rest/api/cosmos-db/common-cosmosdb-rest-request-headers
*/
package rest

import (
	"fmt"
	"net/http"
)

// Header Names
const (
	Authorization   = "Authorization"
	ContentType     = "Content-Type"
	XMsDate         = "x-ms-date"
	XMsSessionToken = "x-ms-session-token"
	XMsVersion      = "x-ms-version"
)

// Error Messages
const (
	errRequiredHeaderMissing = "Required Request Header missing: %s"
)

// Required Headers
// TODO: Determine if XMsSessionToken is always required
func setRequiredHeaders(req *http.Request, headers map[string]string) error {

	if isEmpty(headers[Authorization]) {
		return fmt.Errorf(errRequiredHeaderMissing, Authorization)
	}

	if isEmpty(headers[ContentType]) && contentTypeRequired(req.Method) {
		return fmt.Errorf(errRequiredHeaderMissing, ContentType)
	}

	if isEmpty(headers[XMsDate]) {
		return fmt.Errorf(errRequiredHeaderMissing, XMsDate)
	}

	if isEmpty(headers[XMsVersion]) {
		return fmt.Errorf(errRequiredHeaderMissing, XMsDate)
	}

	// Set All Headers
	for key, val := range headers {
		req.Header.Set(key, val)
	}

	return nil
}

// Content-Type Header is only required on PUT and POST
func contentTypeRequired(method string) bool {
	return (method == http.MethodPut || method == http.MethodPost)
}
