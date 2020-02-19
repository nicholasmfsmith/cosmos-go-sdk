// Package client manages the Azure Cosmos client and provides the functionality to create an instance of a Database.
package client

// Client is the type that describes the Azure Cosmos Client.
type Client struct {
	url        string
	primaryKey string
}

// CosmosClient returns an instance of the Client struct.
func CosmosClient(url, primaryKey string) Client {
	return Client{
		url,
		primaryKey,
	}
}

// TODO: [NS] Create instance of Database
// Database returns an instance of Database with the current instance of Client as the context on which to create the new Database.
// func (client Client) Database(name string) *database.Database {
// 	return &database.Database{
// 		name,
// 	}
// }
