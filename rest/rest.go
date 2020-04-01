// Package rest provides functionality to handle
// the HTTP interactions with the Azure API.
package rest

// Post performs a POST HTTP request to the Azure API to create the provided resource.
// It returns the created resource as a byte array and any errors encountered.
func Post(resource []byte) ([]byte, error) {
	return []byte(""), nil
}

// Get performs a GET HTTP request to the Azure API to read the resource
// identified by the provided resource ID.
// It returns the requested resource as a byte array and any errors encountered.
func Get(id string) ([]byte, error) {
	return []byte(""), nil
}

// Put performs a PUT HTTP request to the Azure API for the provided
// resource ID to replace the existing remote resource with the provided resource.
// It returns the updated resource as a byte array and any errors encountered.
func Put(id string, resource []byte) ([]byte, error) {
	return []byte(""), nil
}

// Delete performs a DELETE HTTP request to the Azure API to remove the resource
// identified by the provided resource ID.
// It returns any errors encountered.
func Delete(id string) error {
	return nil
}
