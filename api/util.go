package api

import "strings"

// Seperate when provided a string simply returns that same string however
// seperated into a slice by newline. We also provide some trimming which is
// why we don't just use the Split function from within the strings package.
func Seperate(s string) []string {
	var blank []string
	for _, s := range strings.Split(s, "\n") {
		blank = append(blank, strings.Trim(s, " "))
	}
	return blank
}
