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

// IToken is the interface to interact with a Token.
type IToken interface {
	Build(method, resourceType, resourcePath, key string) (string, error)
}

// Token defines the required context for an authorization token.
type Token struct {
	method       string
	resourceType string
	resourcePath string
	key          string
	date         string
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
)

// Build generates a new token based on the provided context
// and returns the created authorization string and any errors encountered.
func (token *Token) Build(method, resourceType, resourcePath, key string) (string, error) {
	token.method = method
	token.resourceType = resourceType
	token.resourcePath = resourcePath
	token.key = key
	token.date = strings.ToLower(time.Now().UTC().Format(http.TimeFormat)) // Similar to RFC1123 but uses GMT as the time zone

	decodedKey, err := base64.StdEncoding.DecodeString(token.key)
	if err != nil {
		return "", errors.New(errDecodingProvidedKey + err.Error())
	}

	h := hmac.New(sha256.New, decodedKey)
	text := strings.ToLower(token.method) + "\n" + token.resourceType + "\n" + token.resourcePath + "\n" + token.date + "\n" + "" + "\n"
	_, err = h.Write([]byte(text))
	if err != nil {
		return "", errors.New(errWritingProvidedContextToHash + err.Error())
	}

	sig := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return url.QueryEscape("type=" + masterToken + "&ver=" + tokenVersion + "&sig=" + sig), nil
}
