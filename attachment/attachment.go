/*
Package attachment implements Attachment resource
https://docs.microsoft.com/en-us/rest/api/cosmos-db/attachments
*/
package attachment

// Attachment resource
type Attachment struct {
	name            string
	databaseAccount string
	databaseID      string
	collectionID    string
	documentID      string
	uri             string
	resourceType    string
	partitionKey    string
}

// New returns a new instance of an Attachment resource
func New(partitionKey string) *Attachment {
	return &Attachment{
		uri:          "",
		resourceType: "",
		partitionKey: partitionKey,
	}
}

// URI returns the resource identifier of resource
func (a *Attachment) URI() string {
	return ""
}

// ResourceType returns the resourceType of
func (a *Attachment) ResourceType() string {
	return ""
}

// PartitionKey .
func (a *Attachment) PartitionKey() string {
	return ""
}
