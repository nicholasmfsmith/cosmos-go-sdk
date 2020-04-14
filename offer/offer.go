/*
Package offer implements offer resource type
https://docs.microsoft.com/en-us/rest/api/cosmos-db/offers
*/
package offer

// Offer resource
type Offer struct {
	id              string
	databaseAccount string
}
