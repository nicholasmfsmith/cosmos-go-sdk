// Package container provides functions to create, list, and fetch Cosmos Container resources
// and retrieve and instance of an container client
package container

// Client defines the container client
type Client struct {
	name   string
	dbName string
	key    string
}

// Container defines an cosmos container instance
type Container struct {
	id             string         // name of container provided on creation
	indexingPolicy IndexingPolicy // It is the indexing policy settings for collection.
	partitionKey   PartitionKey   // It is the partitioning configuration settings for collection.
	_rid           string         // It is a system generated property. The resource ID (_rid) is a system-generated identifier.
	_ts            int64          // It is a system generated property. It specifies the last updated timestamp of the resource. The value is a timestamp.
	_self          string         // It is a system generated property. It is the unique addressable URI for the resource.
	_etag          string         // It is a system generated property that specifies the resource etag required for optimistic concurrency control.
	_doc           string         // It is a system generated property that specifies the addressable path of the documents resource.
	_sprocs        string         // It is a system generated property that specifies the addressable path of the stored procedures (sprocs) resource.
	_triggers      string         // It is a system generated property that specifies the addressable path of the triggers resource.
	_udfs          string         // It is a system generated property that specifies the addressable path of the user-defined functions (udfs) resource.
	_conflicts     string         // It is a system generated property that specifies the addressable path of the conflicts resource. During an operation on a resource within a collection, if a conflict occurs, users can inspect the conflicting resources by performing a GET on the conflicts URI path.
}

type PartitionKey struct {
	path {}string
	king string
}

type IndexingPolicy struct {
	automatic     string
	indexingMode  string
	includedPaths IncludedPaths
	excludedPaths ExcludedPaths
}

type IncludedPaths struct {
	path 		string
	dataType 	string // should be enum
	kind 		string
	precision 	int
}

type ExcludedPaths struct {
	path string
}

// Create creates an new instance of an cosmos container
// It returns a Container struct
func (client *Client) Create(id string) *Container {
	return nil
}

// List all available Containers
// It returns a list of Container structs
func (client *Client) List() []*Container {
	return nil
}

// Get fetches a Container by id
// It returns a Container struct
func (client *Client) Get(id string) *Container {
	return nil
}

// Delete deletes an container
// It returns nil
func (client *Client) Delete() {
	return nil
}

// func (containerClient *Client) Item(id, partitionKey string) *item.ItemClient {
// 	return nil
// }
