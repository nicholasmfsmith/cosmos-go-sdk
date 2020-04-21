/*
Package trigger implements offer resource type
https://docs.microsoft.com/en-us/rest/api/cosmos-db/triggers
*/
package trigger

// Trigger resource
type Trigger struct {
	id              string
	databaseAccount string
	uri             string
	resourceType    string
	partitionKey    string
}

// URI returns the resource identifier of resource
func (t *Trigger) URI() string {
	return ""
}

// ResourceType returns the resourceType of
func (t *Trigger) ResourceType() string {
	return ""
}

// PartitionKey .
func (t *Trigger) PartitionKey() string {
	return ""
}
