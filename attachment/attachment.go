/*
Package attachment implements Attachment resource
https://docs.microsoft.com/en-us/rest/api/cosmos-db/attachments
*/
package attachment

// Attachment resource
type Attachment struct {
	name            string
	databaseAccount string
	databaseID      string
	collectionID    string
	documentID      string
}
