// Package client manages the Azure Cosmos client and provides the functionality to create an instance of a Database.
package client

import "cosmos-go-sdk/database"

// Client is the type that describes the Azure Cosmos Client.
type Client struct {
	url string
	key string
}

// New returns an instance of the Client struct.
func New(url, key string) *Client {
	return &Client{
		url,
		key,
	}
}

// Database returns an instance of Database with the current instance of Client as the context on which to create the new Database.
func (client *Client) Database(name string) *database.Database {
	return database.New(name, client.key)
}
