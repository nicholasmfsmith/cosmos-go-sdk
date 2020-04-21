/*
Package sproc implements Stored Procedures resource type
https://docs.microsoft.com/en-us/rest/api/cosmos-db/stored-procedures
*/
package sproc

// StoredProcedure resource
type StoredProcedure struct {
	id              string
	databaseAccount string
	uri             string
	resourceType    string
	partitionKey    string
}

// URI returns the resource identifier of resource
func (s *StoredProcedure) URI() string {
	return ""
}

// ResourceType returns the resourceType of
func (s *StoredProcedure) ResourceType() string {
	return ""
}

// PartitionKey .
func (s *StoredProcedure) PartitionKey() string {
	return ""
}
