package database

import "cosmos-go-sdk/container"

type DatabaseClient struct {
	name string
}

func (databaseClient *DatabaseClient) Container(name string) *container.ContainerClient {
	return nil
}
