// Package item provides CRUD operations for an Azure Cosmos Item.
package item

// Item is the type that describes the Azure Cosmos Item.
type Item struct {
	id            string
	databaseName  string
	containerName string
	key           string
	uri           string
	resourceType  string
	partitionKey  string
}

// URI returns the resource identifier of resource
func (i *Item) URI() string {
	return ""
}

// ResourceType returns the resourceType of
func (i *Item) ResourceType() string {
	return ""
}

// PartitionKey .
func (i *Item) PartitionKey() string {
	return ""
}

/**
TODO: [NS] Remove the interface below. It should be standardized across all sub-resources of container to have a single contract between container and all child resources
type IITem interface {
	// Takes in the document as a []byte and optionally returns error
	Create([]byte) error
	// Returns the document as a []byte and an optional error
	Read() ([]byte, error)
	// Takes in the updated document as a []byte and optionally returns error
	Update([]byte) error
	// Returns an optional error
	Delete() error
}
**/

// New returns an instance of the Item struct.
func New(id, partitionKey, databaseName, containerName, key string) Item {
	return Item{
		id,
		partitionKey,
		databaseName,
		containerName,
		key,
		"", "",
	}
}

// TODO: Consider all SQL Queries are creates
// TODO: Separate HTTP request to utils

// Create creates an Item in the Azure Cosmos Database Container.
// It returns any errors encountered.
func (i *Item) Create(document []byte) error {
	return nil
}

// Read reads an Item in the Azure Cosmos Database Container.
// It returns a byte array of the item in the Azure Cosmos Database Container and any errors encountered.
func (i *Item) Read() ([]byte, error) {
	return []byte(""), nil
}

// Update updates an Item in the Azure Cosmos Database Container.
// It returns any errors encountered.
func (i *Item) Update(document []byte) error {
	return nil
}

// Delete deletes an Item in the Azure Cosmos Database Container.
// It returns any errors encountered.
func (i *Item) Delete() error {
	return nil
}
