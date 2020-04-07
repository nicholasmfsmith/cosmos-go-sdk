/*
Package rest resources source file contains type definitions needed to generate resource-specific requests
*/
package rest

// Resource .
type Resource struct {
	Account    string
	ResourceID string
}

// Database .
type Database struct {
	Resource
}

// User .
type User struct {
	Resource
	DatabaseID string
}

// Permission .
type Permission struct {
	Resource
	UserName   string
	DatabaseID string
}

// Collection .
type Collection struct {
	Resource
	DatabaseID string
}

// StoredProcedure .
type StoredProcedure struct {
	Resource
}

// Trigger .
type Trigger struct {
	Resource
}

// UDF .
type UDF struct {
	Resource
}

// Document .
type Document struct {
	Resource
	DatabaseID   string
	CollectionID string
}

// Attachment .
type Attachment struct {
	Resource
}

// Offer .
type Offer struct {
	Resource
}
