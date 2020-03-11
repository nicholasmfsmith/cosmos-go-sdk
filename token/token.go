package token

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"net/url"
)

const (
	masterToken  = "master"
	tokenVersion = "1.0"
)

type ITokenText interface {
	BuildText(resourceType, resourceID, date string) string
}

type CreateToken string
type ReadToken string
type UpdateToken string
type DeleteToken string

type Token struct {
	ResouceType string
	ResourceID  string
	Date        string
	Token       string
}

func getToken(iTokenText ITokenText) {
	text := iTokenText.BuildText("resourceType", "resourceID", "date")
	decodedKey, _ := base64.StdEncoding.DecodeString("key")

	h := hmac.New(sha256.New, decodedKey)
	h.Write([]byte(text))

	sig := base64.StdEncoding.EncodeToString(h.Sum(nil))
	authToken := url.QueryEscape("type=" + masterToken + "&ver=" + tokenVersion + "&sig=" + sig)
}

func (createToken CreateToken) BuildText(resourceType, resourceID, date string) string {
	return "post" + "\n" + resourceType + "\n" + resourceID + "\n" + date + "\n" + "" + "\n"
}

func (readToken ReadToken) BuildText(resourceType, resourceID, date string) string {
	return "read" + "\n" + resourceType + "\n" + resourceID + "\n" + date + "\n" + "" + "\n"
}

func (updateToken UpdateToken) BuildText(resourceType, resourceID, date string) string {
	return "put" + "\n" + resourceType + "\n" + resourceID + "\n" + date + "\n" + "" + "\n"
}

func (deleteToken DeleteToken) BuildText(resourceType, resourceID, date string) string {
	return "delete" + "\n" + resourceType + "\n" + resourceID + "\n" + date + "\n" + "" + "\n"
}
