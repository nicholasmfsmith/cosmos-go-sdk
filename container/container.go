// Package container provides functions to create, list, and fetch Cosmos Entity resources
// and retrieve and instance of an container client
package container

// Container defines the container client
type Container struct {
	name         string
	dbName       string
	key          string
	uri          string
	resourceType string
	partitionKey string
}

// Client creates an instance of a container
// It returns a Container Client
func Client(name, dbName, key string) *Container {
	return &Container{
		name,
		dbName,
		key,
		"", "", "",
	}
}

// Entity defines an cosmos Container entity
type Entity struct {
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

// PartitionKey defines the partitioning configuration settings for collection.
type PartitionKey struct {
	path []string // An array of paths using which data within the collection can be partitioned.
	kind string   // The algorithm used for partitioning. Only Hash is supported.
}

// IndexingPolicy defines the indexing policy settings for collection.
type IndexingPolicy struct {
	automatic     string          // Indicates whether automatic indexing is on or off. The default value is True, thus all documents are indexed. Setting the value to False would allow manual configuration of indexing paths
	indexingMode  string          // By default, the indexing mode is Consistent. This means that indexing occurs synchronously during insertion, replacment or deletion of documents. To have indexing occur asynchronously, set the indexing mode to lazy.
	includedPaths []IncludedPaths // The array containing document paths to be indexed. By default, two paths are included: the / path, which specifies that all document paths be indexed, and the _ts path, which indexes for a timestamp range comparison.
	excludedPaths []ExcludedPaths // The array containing document paths to be excluded from indexing.
}

// IncludedPaths defines included paths used for indexing
type IncludedPaths struct {
	path      string // Path for which the indexing behavior applies to
	dataType  string // It is the datatype for which the indexing behavior is applied to. Can be String, Number, Point, Polygon, or LineString. Booleans and nulls are automatically indexed
	kind      string // The type of index. Hash indexes are useful for equality comparisons while Range indexes are useful for equality, range comparisons and sorting. Spatial indexes are useful for spatial queries.
	precision int    // The precision of the index. Can be either set to -1 for maximum precision or between 1-8 for Number, and 1-100 for String. Not applicable for Point, Polygon, and LineString data types.
}

// ExcludedPaths defines excluded paths used for indexing
type ExcludedPaths struct {
	path string // Path that is excluded from indexing
}

// Get fetches a Container Entity by id
// It returns a Container Entity struct
func (c *Container) Get() (Entity, error) {
	// TODO - [SC] implement Get
	return Entity{}, nil
}

// Delete deletes an container
// It returns nil if successfull
func (c *Container) Delete() error {
	// TODO - [SC] implement Delete
	return nil
}

// Replace upserts a container to a given database
// It returns a Container Entity struct
func (c *Container) Replace(document Entity) (Entity, error) {
	// TODO - [SC] implement Replace
	return Entity{}, nil
}

// TODO - [SC] need more understanding of what his is suppose to return
// GetPartitionKeyRanges fetches a Entity by id
// It returns a Container Entity struct
// func (client *Container) GetPartitionKeyRanges() *Entity {
// 	return nil
// }

// URI returns the resource identifier of resource
func (c *Container) URI() string {
	return ""
}

// ResourceType returns the resourceType of
func (c *Container) ResourceType() string {
	return ""
}

// PartitionKey .
func (c *Container) PartitionKey() string {
	return ""
}
