package container

import (
	"cosmos-go-sdk/item"
)

type Container struct {
	name string
}

// TODO: [NS] Create an interface for all resources that are valid at the container level
// TODO: [NS] Create an 'Items' package to handle actions that are performed across items
func (container *Container) Item(id, partitionKey string) *item.Item {
	return nil
}
