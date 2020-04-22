// Package document provides CRUD operations for an Azure Cosmos document.
package document

import "cosmos-go-sdk/rest"

// Document is the type that describes the Azure Cosmos document.
type Document struct {
	id           string
	partitionKey string
	uri          string
	key          string
}

// New returns an instance of the document struct.
func New(id, partitionKey, containerURI, key string) Document {
	return Document{
		id,
		partitionKey,
		containerURI + "/docs/" + id,
		key,
	}
}

// TODO: Consider all SQL Queries are creates
// TODO: Separate HTTP request to utils

// Create creates an document in the Azure Cosmos Database Container.
// It returns any errors encountered.
func (document Document) Create(doc []byte) error {
	return nil
}

// Read reads a document in the Azure Cosmos Database Container.
// It returns a byte array of the document in the Azure Cosmos Database Container and any errors encountered.
func (document Document) Read() ([]byte, error) {
	doc, err := rest.Get()
	if err != nil {
		return nil, err
	}
	return doc, nil
}

// Update updates an document in the Azure Cosmos Database Container.
// It returns any errors encountered.
func (document Document) Update(doc []byte) error {
	return nil
}

// Delete deletes an document in the Azure Cosmos Database Container.
// It returns any errors encountered.
func (document Document) Delete() error {
	return nil
}

func (document Document) URI() string {
	return document.uri
}
func (document Document) ResourceType() string {
	return "docs"

}
func (document Document) PartitionKey() string {
	return document.partitionKey
}
