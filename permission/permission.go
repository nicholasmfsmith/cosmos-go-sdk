/*
Package permission implements permission resource type
https://docs.microsoft.com/en-us/rest/api/cosmos-db/permissions
*/
package permission

// Permission resource
type Permission struct {
	id              string
	databaseAccount string
	userName        string
	databaseID      string
	uri             string
	resourceType    string
	partitionKey    string
}

// URI returns the resource identifier of resource
func (p *Permission) URI() string {
	return ""
}

// ResourceType returns the resourceType of
func (p *Permission) ResourceType() string {
	return ""
}

// PartitionKey .
func (p *Permission) PartitionKey() string {
	return ""
}
