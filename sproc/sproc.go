/*
Package sproc implements Stored Procedures resource type
https://docs.microsoft.com/en-us/rest/api/cosmos-db/stored-procedures
*/
package sproc

// StoredProcedure resource
type StoredProcedure struct {
	id              string
	databaseAccount string
}
