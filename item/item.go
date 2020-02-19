package item

type ItemClient struct {
	id           string
	partitionKey string
}

// TODO: Consider all SQL Queries are creates
// TODO: Separate HTTP request to utils
// TODO: Explore potentially using async logic for HTTP requests
func (itemClient *ItemClient) Create() error {
	return nil
}

// TODO: Look into the differences in CRUD for different Cosmos components (it's possible the methods are similar, may be able to generalize)
func (itemClient *ItemClient) Read() ([]byte, error) {
	return nil, nil
}
