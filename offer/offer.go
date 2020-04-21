/*
Package offer implements offer resource type
https://docs.microsoft.com/en-us/rest/api/cosmos-db/offers
*/
package offer

// Offer resource
type Offer struct {
	id              string
	databaseAccount string
	uri             string
	resourceType    string
	partitionKey    string
}

// New returns a new instance of an Attachment resource
func New(partitionKey string) *Offer {
	return &Offer{
		uri:          "",
		resourceType: "",
		partitionKey: partitionKey,
	}
}

// URI returns the resource identifier of resource
func (o *Offer) URI() string {
	return ""
}

// ResourceType returns the resourceType of
func (o *Offer) ResourceType() string {
	return ""
}

// PartitionKey .
func (o *Offer) PartitionKey() string {
	return ""
}
