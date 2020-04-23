// Package document provides CRUD operations for an Azure Cosmos document.
package document

import "cosmos-go-sdk/rest"

// Document is the type that describes the Azure Cosmos document.
type Document struct {
	ID           string
	PartitionKey string
	URI          string
	Key          string
	Request      rest.IRequest
}

// New returns an instance of the document struct.
func New(id, partitionKey, containerURI, key string) Document {
	uri := containerURI + "/docs/" + id
	return Document{
		id,
		partitionKey,
		uri,
		key,
		// TODO: [NS] Create util function for building URI
		rest.New(uri, "docs", key),
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
	doc, err := document.Request.Get()
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
