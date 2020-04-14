/*
Package users implements user resource type
https://docs.microsoft.com/en-us/rest/api/cosmos-db/users
*/
package users

// User resource
type User struct {
	id              string
	databaseAccount string
	databaseID      string
}
