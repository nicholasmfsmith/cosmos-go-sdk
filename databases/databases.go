package databases

import "cosmos-go-sdk/database"

// Databases defines the Databases Client used to
// perform actions
type Databases struct {
	name string
	key  string
}

// Client returns an instance of the Database struct.
func Client(name, key string) Databases {
	return Databases{
		name,
		key,
	}
}

// Create creates an new instance of an cosmos database
// It returns a Database entity
func (client *Databases) Create(id string) database.Entity {
	return nil
}

// CreateIfNotExist creates an new instance of an cosmos database
// It returns a Database entity
func (client *Databases) CreateIfNotExist(id string) database.Entity {
	return nil
}

// List queries all databases on an given account
// It should return an array of Entities
func (client *Databases) List() []database.Entity {
	return nil
}
