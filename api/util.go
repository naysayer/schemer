package api

import "strings"

const (
	//TypeInt represents a golang int data type as a string
	TypeInt = "int"
	//TypeFloat represents a golang float data type as a string
	TypeFloat = "float64"
	//TypeString represents a golang string data type as a string
	TypeString = "string"
	//TypeTime represents a golang time.Time data type as a string
	TypeTime = "time.Time"
	//TypeJSONText represents a golang types.JSONText data type as a string, this stems from within sqlx
	TypeJSONText = "types.JSONText"
	//TypeBool represents a golang bool data type as a string
	TypeBool = "bool"

	// DbStructTag is a tag that tha tis used to point to what the colums name is in the db
	DbStructTag = "`db:\"%v\"`"
)

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
