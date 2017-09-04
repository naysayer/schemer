// Package structure used to define a structure interface and methods around it
// that are ultimately used for the generation of golang structs
package structure

import (
	"bytes"

	"github.com/naysayer/schemer/api/structure/attr"
)

// Structure is an interface that is used to represent a block of text
// within a schema file that ultimate represents a database tabe. Where
// title representes the title of the table, and contents, (a slice of attrs)
// would be a list of the corresponding columns within said table.
type Structure interface {
	Title() string
	Contents() []attr.Attr
}

// Stringify accepting a structure interface returns a string that represents
// the native golang equilivent of that structure as a golang struct.
func Stringify(s Structure) string {
	var buf bytes.Buffer

	buf.WriteString("type ")
	buf.WriteString(s.Title())
	buf.WriteString(" struct {")

	for _, attr := range s.Contents() {
		buf.WriteString("\n")
		buf.WriteString(attr.Title())
		buf.WriteString(" ")
		buf.WriteString(attr.Type())
		buf.WriteString(" ")
		buf.WriteString(attr.Tags())
	}

	buf.WriteString("\n")
	buf.WriteString("}")

	return buf.String()
}
