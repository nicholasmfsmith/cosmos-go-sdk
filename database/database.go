// Package database provides functions to create, list, and fetch Cosmos Database resources
// and retrieve and instance of an container client
package database

// Database defines the database client
type Database struct {
	name string
	key  string
}

// Client returns an instance of the Client struct.
func Client(name string, key string) Database {
	return Database{
		name,
		key,
	}
}

// Document defines an cosmos database instance
type Document struct {
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
func (client *Database) Create(id string) *Document {
	return nil
}

// List all available Databases
// It returns a list of Database structs
func (client *Database) List() ([]*Document, error) {
	return []*Document{&Document{}}, nil
}

// Get fetches a Database by id
// It returns a Database struct
func (client *Database) Get(id string) (*Document, error) {
	return &Document{}, nil
}

// Delete an database
func (client *Database) Delete(id string) (*Document, error) {
	return &Document{}, nil
}

// Container returns an container client
// func (client *Database) Container(name string) *container.Client {
// 	return nil
// }
