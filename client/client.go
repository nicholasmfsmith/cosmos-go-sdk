package client

import "cosmos-go-sdk/database"

type BaseClient struct {
	url string
	key string
}

func CosmosClient(url, key string) *BaseClient {
	return &BaseClient{
		url,
		key,
	}
}

func (baseClient *BaseClient) Database(name string) *database.DatabaseClient {
	return nil
}
