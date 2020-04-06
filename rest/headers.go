/*
Package rest headers source file addresses the request headers used per spec  https://docs.microsoft.com/en-us/rest/api/cosmos-db/common-cosmosdb-rest-request-headers
TODO: Handling of Optional Headers; Currently only Required Headers are being accounted for.
*/
package rest

import (
	"errors"
	"net/http"
	"strings"
)

// Header Names
const (
	headerAuthorization   = "Authorization"
	headerContentType     = "Content-Type"
	headerXMsDate         = "x-ms-date"
	headerXMsSessionToken = "x-ms-session-token"
	headerXMsVersion      = "x-ms-version"
)

// Headers contains common request headers with SQL API
type Headers struct {
	Authorization   string
	ContentType     string
	XMsDate         string // rfc1123
	XMsSessionToken string
	XMsVersion      string
}

// Header Error Messages
const (
	errNoAuthorization   = "Required Request Header missing: Authorization"
	errNoContentType     = "Required Request Header missing: Content-Type (on PUT and POST)"
	errNoXMsDate         = "Required Request Header missing: x-ms-date"
	errNoXMsSessionToken = "Required Request Header missing: x-ms-session-token (for session consistency only)"
	errNoXMsVersion      = "Required Request Header missing: x-ms-version"
)

// Required Headers
// Authorization, Content-Type (only Required on PUT and POST), XMsDate, XMsSessionToken (for session consistency only), XMsVersion
// TODO: Determine if required validator is necessary XMsSessionToken
func setRequiredHeaders(req *http.Request, authorization, contentType, xMsDate, xMsVersion string) error {

	if emptyString(authorization) {
		return errors.New(errNoAuthorization)
	}

	if !emptyString(contentType) {
		req.Header.Set(headerContentType, contentType)
	} else if contentTypeRequired(req.Method) {
		return errors.New(errNoContentType)
	}

	if emptyString(xMsDate) {
		return errors.New(errNoXMsDate)
	}

	if emptyString(xMsVersion) {
		return errors.New(errNoXMsVersion)
	}

	req.Header.Set(headerAuthorization, authorization)
	req.Header.Set(headerXMsDate, xMsDate)
	req.Header.Set(headerXMsVersion, xMsVersion)

	return nil
}

// Content-Type Header is only required on PUT and POST
func contentTypeRequired(method string) bool {
	return (method == http.MethodPut || method == http.MethodPost)
}

// Trims white-space then checks length of string
func emptyString(val string) bool {
	return len(strings.TrimSpace(val)) == 0
}

// TODO: Account for Optional Headers
/*

const (
	headerAuthorization                          = "Authorization"
	headerContentType                            = "Content-Type"
	headerIfMatch                                = "If-Match"
	headerIfNoneMatch                            = "If-None-Match"
	headerIfModifiedSince                        = "If-Modified-Since"
	headerUserAgent                              = "User-Agent"
	headerXMsActivityID                          = "x-ms-activity-id"
	headerXMsConsistencyLevel                    = "x-ms-consistency-level"
	headerXMsContinuation                        = "x-ms-continuation"
	headerXMsDate                                = "x-ms-date"
	headerXMsMaxItemCount                        = "x-ms-max-item-count"
	headerXMsDocumentDBPartitionKey              = "x-ms-documentdb-partitionkey"
	headerXMsDocumentDBQueryEnableCrossPartition = "x-ms-documentdb-query-enablecrosspartition"
	headerXMsSessionToken                        = "x-ms-session-token"
	headerXMsVersion                             = "x-ms-version"
	headerAIm                                    = "A-IM"
	headerXMsDocumentDBPartitionKeyRangeID       = "x-ms-documentdb-partitionkeyrangeid"
	headerXMsCosmosAllowTentativeWrites          = "x-ms-cosmos-allow-tentative-writes"
)

type Headers struct {
	Authorization                          string
	ContentType                            string
	IfMatch                                string
	IfNoneMatch                            string
	IfModifiedSince                        string // rfc1123
	UserAgent                              string
	XMsActivityID                          string
	XMsConsistencyLevel                    string
	XMsContinuation                        string
	XMsDate                                string // rfc1123
	XMsMaxItemCount                        string
	XMsDocumentDBPartitionKey              string
	XMsDocumentDBQueryEnableCrossPartition string
	XMsSessionToken                        string
	XMsVersion                             string
	AIm                                    string
	XMsDocumentDBPartitionKeyRangeID       string
	XMsCosmosAllowTentativeWrites          string
}

// TODO: Rework with a map for optional headers
func setHeaders(req *http.Request, headers Headers) error {

	errRequiredHeaders := setRequiredHeaders(req, headers.Authorization, headers.ContentType, headers.XMsDate, headers.XMsVersion)
	if errRequiredHeaders != nil {
		return errRequiredHeaders
	}

	// Common, optional headers
	if !emptyString(headers.IfMatch) {
		req.Header.Set(headerIfMatch, headers.IfMatch)
	}
	if !emptyString(headers.IfNoneMatch) {
		req.Header.Set(headerIfNoneMatch, headers.IfNoneMatch)
	}
	if !emptyString(headers.UserAgent) {
		req.Header.Set(headerUserAgent, headers.UserAgent)
	}
	if !emptyString(headers.XMsActivityID) {
		req.Header.Set(headerXMsActivityID, headers.XMsActivityID)
	}
	if !emptyString(headers.XMsConsistencyLevel) {
		req.Header.Set(headerXMsConsistencyLevel, headers.XMsConsistencyLevel)
	}
	if !emptyString(headers.XMsContinuation) {
		req.Header.Set(headerXMsContinuation, headers.XMsContinuation)
	}
	if !emptyString(headers.XMsMaxItemCount) {
		req.Header.Set(headerXMsMaxItemCount, headers.XMsMaxItemCount)
	}
	if !emptyString(headers.XMsDocumentDBPartitionKey) {
		req.Header.Set(headerXMsDocumentDBPartitionKey, headers.XMsDocumentDBPartitionKey)
	}
	if !emptyString(headers.XMsDocumentDBQueryEnableCrossPartition) {
		req.Header.Set(headerXMsDocumentDBQueryEnableCrossPartition, headers.XMsDocumentDBQueryEnableCrossPartition)
	}
	if !emptyString(headers.XMsSessionToken) {
		req.Header.Set(headerXMsSessionToken, headers.XMsSessionToken)
	}
	if !emptyString(headers.AIm) {
		req.Header.Set(headerAIm, headers.AIm)
	}
	if !emptyString(headers.XMsDocumentDBPartitionKeyRangeID) {
		req.Header.Set(headerXMsDocumentDBPartitionKeyRangeID, headers.XMsDocumentDBPartitionKeyRangeID)
	}
	if !emptyString(headers.XMsCosmosAllowTentativeWrites) {
		req.Header.Set(headerXMsCosmosAllowTentativeWrites, headers.XMsCosmosAllowTentativeWrites)
	}

	return nil
}
*/
