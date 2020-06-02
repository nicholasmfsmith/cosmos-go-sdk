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

// IDocument is the interface that defines the functionality of the document package
type IDocument interface {
	Read() ([]byte, error)
	Create(doc []byte) ([]byte, error)
	Update(doc []byte) ([]byte, error)
	Delete() error
}

// New returns an instance of the document struct.
func New(id, partitionKey, collectionURI, key string) Document {
	// TODO: ID & PartitionKey empty string check and handling
	// TODO: [NS] Create util function for building URI
	uri := collectionURI + "/docs/" + id
	return Document{
		id,
		partitionKey,
		uri,
		key,
		rest.New(uri, "docs", key),
	}
}

// Create creates a document in the Azure Cosmos Database Collection.
// It returns any errors encountered.
func (document Document) Create(doc []byte) ([]byte, error) {
	doc, err := document.Request.Post(doc)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

// Read reads a document in the Azure Cosmos Database Collection.
// It returns a byte array of the document in the Azure Cosmos Database Collection and any errors encountered.
func (document Document) Read() ([]byte, error) {
	doc, err := document.Request.Get()
	if err != nil {
		return nil, err
	}
	return doc, nil
}

// Update a document in the Azure Cosmos Database Collection.
// It returns any errors encountered.
func (document Document) Update(doc []byte) ([]byte, error) {
	doc, err := document.Request.Put(doc)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

// Delete deletes an document in the Azure Cosmos Database Collection.
// It returns any errors encountered.
func (document Document) Delete() error {
	return nil
}
