/*
Package rest uris source file addresses the Resource URI Syntax spec
https://docs.microsoft.com/en-us/rest/api/cosmos-db/cosmosdb-resource-uri-syntax-for-rest
*/
package rest

import (
	"fmt"
)

const (
	errMissingParameter    = "Resource Parameter Missing: %s"
	errNoGETMethod         = "Error building GET URI: Cannot perform GET on resource:"
	errInvalidResourceType = "Error building GET URI: Invalid Resource Type;"
)

/*
buildGETURI creates GET request URI
Base URI for a resource is composed of: DB account name,  List of resource types, names of req resource and its parents
	- Database		https://{databaseaccount}.documents.azure.com/dbs/{db-id}
	- User				https://{databaseaccount}.documents.azure.com/dbs/{db-id}/users/{user-name}
	- Permission	https://{databaseaccount}.documents.azure.com/dbs/{db-id}/users/{user-name}/permissions/{permission-name}
	- Collection	https://{databaseaccount}.documents.azure.com/dbs/{db-id}/colls/{coll-id}
	- Document		https://{databaseaccount}.documents.azure.com/dbs/{db-id}/colls/{coll-id}/docs/{doc-id}
	- Offer				https://{databaseaccount}.documents.azure.com/offers/{_rid-offer}
*/
func createGETURI(resource interface{}) (string, error) {
	var uri string
	var errURI error

	switch resourceType := resource.(type) {
	case Database:
		uri, errURI = createURI("https://%s.documents.azure.com/dbs/%s",
			resource.(Database).Account, resource.(Database).ResourceID)
	case User:
		uri, errURI = createURI("https://%s.documents.azure.com/dbs/%s/users/%s",
			resource.(User).Account, resource.(User).DatabaseID,
			resource.(User).ResourceID)
	case Permission:
		uri, errURI = createURI("https://%s.documents.azure.com/dbs/%s/users/%s/permissions/%s",
			resource.(Permission).Account, resource.(Permission).DatabaseID,
			resource.(Permission).UserName, resource.(Permission).ResourceID)
	case Collection:
		uri, errURI = createURI("https://%s.documents.azure.com/dbs/%s/colls/%s",
			resource.(Collection).Account, resource.(Collection).DatabaseID,
			resource.(Collection).ResourceID)
	case Document:
		uri, errURI = createURI("https://%s.documents.azure.com/dbs/%s/colls/%s/docs/%s",
			resource.(Document).Account, resource.(Document).DatabaseID,
			resource.(Document).CollectionID, resource.(Document).ResourceID)
	case Offer:
		uri, errURI = createURI("https://%s.documents.azure.com/offers/%s",
			resource.(Offer).Account, resource.(Offer).ResourceID)
	case StoredProcedure, Trigger, UDF, Attachment:
		return "", fmt.Errorf("%s got %s", errNoGETMethod, resourceType)
	default:
		return "", fmt.Errorf("%s got %s", errInvalidResourceType, resourceType)
	}

	if errURI != nil {
		return "", errURI
	}

	return uri, nil
}

// createURI constructs URI and checks for nil params
func createURI(uri string, args ...interface{}) (string, error) {

	// Check for nil arguments
	for _, val := range args {
		str := fmt.Sprintf("%v", val)
		if isEmpty(str) {
			return "", fmt.Errorf(errMissingParameter, val)
		}
	}

	// Pass arguments
	uri = fmt.Sprintf(uri, args...)
	return uri, nil
}
