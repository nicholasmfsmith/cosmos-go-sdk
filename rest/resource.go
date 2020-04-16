package rest

// IResource defines a CosmosDB REST API resource
type IResource interface {
	URI() string
	ResourceType() string
	PartitionKey() string
	ID() string
}
