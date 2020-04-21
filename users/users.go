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
	uri             string
	resourceType    string
	partitionKey    string
}

// URI returns the resource identifier of resource
func (u *User) URI() string {
	return ""
}

// ResourceType returns the resourceType of
func (u *User) ResourceType() string {
	return ""
}

// PartitionKey .
func (u *User) PartitionKey() string {
	return ""
}
