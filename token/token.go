package token

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"net/http"
	"net/url"
	"strings"
	"sync"
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
	mux          sync.Mutex
}

// Private constants used to build the token.
const (
	masterToken  = "master"
	tokenVersion = "1.0"
)

// Useful HTTP verbs for the client to leverage.
const (
	MethodGet    = "get"
	MethodPost   = "post"
	MethodPut    = "put"
	MethodDelete = "delete"
)

// Internal error messages used for building token.
const (
	errorDecodingProvidedKey          = "Error decoding provided key: "
	errorWritingProvidedContextToHash = "Error writing provided context to the hash: "
	emptyHTTPMethod                   = "Error building token: Provided HTTP method is empty"
	emptyResourceType                 = "Error building token: Provided Resource Type is empty"
	emptyResourceID                   = "Error building token: Provided Resource ID is empty"
	emptyKey                          = "Error building token: Provided key is empty"
	emptyDate                         = "Error building token: Provided date is empty"
)

// New returns a pointer to a new instance of Token
// based on the provided input context of method, resource type, resource ID, and key.
func New(method, resourceType, resourceID, key string) *Token {
	token := Token{}
	token.Method = method
	token.ResourceType = resourceType
	token.ResourceID = resourceID
	token.Key = key
	// Similar to RFC1123 but uses GMT as the time zone
	token.Date = strings.ToLower(time.Now().UTC().Format(http.TimeFormat))
	return &token
}

// Build generates a new token based on the input context
// provided by the function receiver of *Token and returns any errors encountered.
func (token *Token) Build() error {
	// Error checking for empty context values
	if len(token.Method) == 0 {
		return errors.New(emptyHTTPMethod)
	}

	if len(token.ResourceType) == 0 {
		return errors.New(emptyResourceType)
	}

	if len(token.ResourceID) == 0 {
		return errors.New(emptyResourceID)
	}

	if len(token.Key) == 0 {
		return errors.New(emptyKey)
	}

	if len(token.Date) == 0 {
		return errors.New(emptyDate)
	}

	// Set a mutual exclusion to protect against conflicts between goroutines.
	token.mux.Lock()
	defer token.mux.Unlock()

	decodedKey, err := base64.StdEncoding.DecodeString(token.Key)
	if err != nil {
		return errors.New(errorDecodingProvidedKey + err.Error())
	}

	h := hmac.New(sha256.New, decodedKey)
	text := strings.ToLower(token.Method) + "\n" + token.ResourceType + "\n" + token.ResourceID + "\n" + token.Date + "\n" + "" + "\n"
	_, err = h.Write([]byte(text))
	if err != nil {
		return errors.New(errorWritingProvidedContextToHash + err.Error())
	}

	sig := base64.StdEncoding.EncodeToString(h.Sum(nil))
	token.Token = url.QueryEscape("type=" + masterToken + "&ver=" + tokenVersion + "&sig=" + sig)
	return nil
}
