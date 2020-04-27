// Package database provides functions to create, list, and fetch Cosmos Database resources
// and retrieve and instance of an container client
package database

import (
	"cosmos-go-sdk/rest"
	"encoding/json"
)

const (
	ResourceType = "dbs"
)

// Database defines the database client
type Database struct {
	URI     string
	Name    string
	Key     string
	Request rest.IRequest
}

// New returns an instance of Database
func New(name, key, uri string) Database {
	uri = uri + "/" + ResourceType + "/" + name
	return Database{
		uri,
		name,
		key,
		rest.New(uri, ResourceType, key),
	}
}

// Entity defines an cosmos database entity
type Entity struct {
	ID    string `json:"id"`     // name of database provided on creation
	RID   string `json:"_rid"`   // It is a system generated property. The resource ID (_rid) is a system-generated identifier.
	TS    int64  `json:"_ts"`    // It is a system generated property. It specifies the last updated timestamp of the resource. The value is a timestamp.
	SELF  string `json:"_self"`  // It is a system generated property. It is the unique addressable URI for the resource.
	ETAG  string `json:"_etag"`  // It is a system generated property that specifies the resource etag required for optimistic concurrency control.
	COLLS string `json:"_colls"` // It is a system generated property that specifies the addressable path of the collections resource.
	USER  string `json:"_user"`  // It is a system generated property that specifies the addressable path of the users resource.
}

// Get fetches a Database by id
// It returns a Database entity
func (database Database) Read() (Entity, error) {
	bytes, httpError := database.Request.Get()
	if httpError != nil {
		return Entity{}, httpError
	}
	var databaseEntity Entity
	json.Unmarshal(bytes, &databaseEntity)
	return databaseEntity, nil
}

// Delete an Database entity
// It returns deleted database entity
func (database Database) Delete() error {
	return nil
}
