
package rest

import "strings"

// emptyString Trims white-space then checks length of string
// returns true if string argument is empty
func isEmpty(val string) bool {
	return len(strings.TrimSpace(val)) == 0
}
