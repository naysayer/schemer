package structure

import (
	"testing"

	"github.com/naysayer/schemer/api/structure/attr"
	"github.com/stretchr/testify/assert"
)

type str struct {
}

func (s str) Title() string {
	return "StructName"
}
func (s str) Contents() []attr.Attr {
	return []attr.Attr{
		att{},
	}
}

type att struct {
}

func (a att) Title() string {
	return "AttrName"
}
func (a att) Type() string {
	return "string"
}
func (a att) Tags() string {
	return "`db:\"db_name\"`"
}

func TestStringify(t *testing.T) {
	actual := Stringify(str{})
	expected := "type StructName struct {\nAttrName string `db:\"db_name\"`\n}"

	assert.Equal(t, actual, expected)
}
