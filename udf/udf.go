/*
Package udf implements offer resource type
https://docs.microsoft.com/en-us/rest/api/cosmos-db/user-defined-functions
*/
package udf

// UDF resource
type UDF struct {
	id              string
	databaseAccount string
}
