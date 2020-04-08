/*
Package rest resources source file contains type definitions needed to generate resource-specific requests
TODO: Is there a better place for these definitions? Can they be integrated with resource packages?
*/
package rest

// Resource defines a CosmosDB REST API resource
type Resource struct {
	Account    string
	ResourceID string
}

// Database resource
type Database struct {
	Resource
}

// User resource
type User struct {
	Resource
	DatabaseID string
}

// Permission resource
type Permission struct {
	Resource
	UserName   string
	DatabaseID string
}

// Collection resource
type Collection struct {
	Resource
	DatabaseID string
}

// StoredProcedure resource
type StoredProcedure struct {
	Resource
}

// Trigger resource
type Trigger struct {
	Resource
}

// UDF resource
type UDF struct {
	Resource
}

// Document resource
type Document struct {
	Resource
	DatabaseID   string
	CollectionID string
}

// Attachment resource
type Attachment struct {
	Resource
}

// Offer resource
type Offer struct {
	Resource
}
