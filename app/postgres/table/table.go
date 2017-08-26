package table

import (
	"regexp"
	"strings"
)

// Name: We know because of regex that there will be a value at the index 0 of slice
// of strings representing a create table statement. The line of text at index 0
// is the create table statement, from this line we returned a sanitized create
// table statement.
func Name(s []string) string {
	r := regexp.MustCompile("(CREATE TABLE) | (\\()")
	return strings.Trim(r.ReplaceAllString(s[0], ""), " ")
}
