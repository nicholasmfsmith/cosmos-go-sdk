/*
Package trigger implements offer resource type
https://docs.microsoft.com/en-us/rest/api/cosmos-db/triggers
*/
package trigger

// Trigger resource
type Trigger struct {
	id              string
	databaseAccount string
}
