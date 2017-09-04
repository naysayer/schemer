package table

import (
	"regexp"
	"strings"
)

var (
	patternCreateTableName = regexp.MustCompile(`(CREATE TABLE) | (\()`)
)

// Name returns a string that is represents the name of a postgres database
// table. This returns a sanitized version of the name with leading and trailing
// space stripped away. What is important to note is that this simply strips
// away the text around the name of the table. Rather than extracting it.
// In the event that the pattern is not present then we return the inputted
// string
func Name(s string) string {
	return strings.Trim(patternCreateTableName.ReplaceAllString(s, ""), " ")
}
