// Package token provides functionality to generate an Azure Cosmos token.
package token

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Token is the type describes a constructed token.
// It includes the dynamic context required to create a valid token.
type Token struct {
	Method       string
	ResourceType string
	ResourceID   string
	Key          string
	Date         string
	Token        string
}

// Private constants used to build the token.
const (
	masterToken  = "master"
	tokenVersion = "1.0"
)

// Internal error messages used for building token.
const (
	errDecodingProvidedKey          = "Error decoding provided key: "
	errWritingProvidedContextToHash = "Error writing provided context to the hash: "
	errEmptyHTTPMethod              = "Error building token: Provided HTTP method is empty"
	errEmptyResourceType            = "Error building token: Provided Resource Type is empty"
	errEmptyResourceID              = "Error building token: Provided Resource ID is empty"
	errEmptyKey                     = "Error building token: Provided key is empty"
	errEmptyDate                    = "Error building token: Provided date is empty"
)

// New returns a pointer to a new instance of Token
// based on the provided input context of method, resource type, resource ID, and key.
func New(method, resourceType, resourceID, key string) *Token {
	return &Token{
		Method:       method,
		ResourceType: resourceType,
		ResourceID:   resourceID,
		Key:          key,
		Date:         strings.ToLower(time.Now().UTC().Format(http.TimeFormat)), // Similar to RFC1123 but uses GMT as the time zone
	}
}

// Build generates a new token based on the input context
// provided by the function receiver of *Token and returns any errors encountered.
func (token *Token) Build() error {
	// Error checking for empty context values
	if len(token.Method) == 0 {
		return errors.New(errEmptyHTTPMethod)
	}

	if len(token.ResourceType) == 0 {
		return errors.New(errEmptyResourceType)
	}

	if len(token.ResourceID) == 0 {
		return errors.New(errEmptyResourceID)
	}

	if len(token.Key) == 0 {
		return errors.New(errEmptyKey)
	}

	if len(token.Date) == 0 {
		return errors.New(errEmptyDate)
	}

	decodedKey, err := base64.StdEncoding.DecodeString(token.Key)
	if err != nil {
		return errors.New(errDecodingProvidedKey + err.Error())
	}

	h := hmac.New(sha256.New, decodedKey)
	text := strings.ToLower(token.Method) + "\n" + token.ResourceType + "\n" + token.ResourceID + "\n" + token.Date + "\n" + "" + "\n"
	_, err = h.Write([]byte(text))
	if err != nil {
		return errors.New(errWritingProvidedContextToHash + err.Error())
	}

	sig := base64.StdEncoding.EncodeToString(h.Sum(nil))
	token.Token = url.QueryEscape("type=" + masterToken + "&ver=" + tokenVersion + "&sig=" + sig)
	return nil
}
