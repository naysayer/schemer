package structure

import (
	"bytes"

	"github.com/naysayer/schemer/api/structure/attr"
)

type Structure interface {
	Title() string
	Contents() []attr.Attr
}

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
