// Package database provides functions to create, list, and fetch Cosmos Database resources
// and retrieve and instance of an container client
package database

// Database defines the database client
type Database struct {
	name string
	key  string
}

// Client returns an instance of the Client struct.
func Client(name, key string) Database {
	return Database{
		name,
		key,
	}
}

// Entity defines an cosmos database entity
type Entity struct {
	id     string // name of database provided on creation
	_rid   string // It is a system generated property. The resource ID (_rid) is a system-generated identifier.
	_ts    int64  // It is a system generated property. It specifies the last updated timestamp of the resource. The value is a timestamp.
	_self  string // It is a system generated property. It is the unique addressable URI for the resource.
	_etag  string // It is a system generated property that specifies the resource etag required for optimistic concurrency control.
	_colls string // It is a system generated property that specifies the addressable path of the collections resource.
	_user  string // It is a system generated property that specifies the addressable path of the users resource.
}

// Get fetches a Database by id
// It returns a Database entity
func (client *Database) Get(id string) (Entity, error) {
	return Entity{}, nil
}

// Delete an Database entity
// It returns deleted database entity
func (client *Database) Delete() (Entity, error) {
	return Entity{}, nil
}
