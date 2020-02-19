package container

import (
	"cosmos-go-sdk/item"
)

type ContainerClient struct {
	name string
}

func (containerClient *ContainerClient) Item(id, partitionKey string) *item.ItemClient {
	return nil
}
