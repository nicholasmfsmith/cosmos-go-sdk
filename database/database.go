// Package database provides functions to create, list, and fetch Cosmos Database resources
// and retrieve and instance of an container client
package database

// Client defines the database client
type Client struct {
	name string
	key  string
}

// Database defines an cosmos database instance
type Database struct {
	id     string // name of database provided on creation
	_rid   string // It is a system generated property. The resource ID (_rid) is a system-generated identifier.
	_ts    int64  // It is a system generated property. It specifies the last updated timestamp of the resource. The value is a timestamp.
	_self  string // It is a system generated property. It is the unique addressable URI for the resource.
	_etag  string // It is a system generated property that specifies the resource etag required for optimistic concurrency control.
	_colls string // It is a system generated property that specifies the addressable path of the collections resource.
	_user  string // It is a system generated property that specifies the addressable path of the users resource.
}

// Create creates an new instance of an cosmos database
// It returns a Database struct
func (client *Client) Create(id string) *Database {
	return nil
}

// List all available Databases
// It returns a list of Database structs
func (client *Client) List() []*Database {
	return nil
}

// Get fetches a Database by id
// It returns a Database struct
func (client *Client) Get(id string) *Database {
	return nil
}

// Delete an database
func (client *Client) Delete() {
}

// Container returns an container client
// func (client *Client) Container(name string) *container.Client {
// 	return nil
// }
