/*
Package udf implements offer resource type
https://docs.microsoft.com/en-us/rest/api/cosmos-db/user-defined-functions
*/
package udf

// UDF resource
type UDF struct {
	id              string
	databaseAccount string
	uri             string
	resourceType    string
	partitionKey    string
}

// URI returns the resource identifier of resource
func (u *UDF) URI() string {
	return ""
}

// ResourceType returns the resourceType of
func (u *UDF) ResourceType() string {
	return ""
}

// PartitionKey .
func (u *UDF) PartitionKey() string {
	return ""
}
