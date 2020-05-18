// Package client manages the Azure Cosmos client and provides the functionality to create an instance of a Database.
package client

import "cosmos-go-sdk/database"

// Client is the type that describes the Azure Cosmos Client.
type Client struct {
	uri string
	key string
}

// New returns an instance of the Client struct.
func New(uri, key string) Client {
	return Client{
		uri,
		key,
	}
}

// Database returns an instance of Database with the current instance of Client as the context on which to create the new Database.
func (client Client) Database(name string) database.Database {
	return database.New(name, client.key, client.uri)
}
