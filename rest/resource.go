package rest

// IResource defines a CosmosDB REST API resource
type IResource interface {
	ID() string
	Account() string
	URI() (string, error)
}
