package structure

import (
	"testing"

	"github.com/naysayer/schemer/api/structure/attr"
	"github.com/stretchr/testify/assert"
)

type Str struct {
}

func (s Str) Name() string {
	return "StructName"
}
func (s Str) Contents() []attr.Attr {
	return []attr.Attr{
		Att{},
	}
}

type Att struct {
}

func (a Att) Name() string {
	return "AttrName"
}
func (a Att) Type() string {
	return "string"
}
func (a Att) Tags() string {
	return "`db:\"db_name\"`"
}

func TestStringify(t *testing.T) {
	actual := Stringify(Str{})
	expected := "type StructName struct {\nAttrName string `db:\"db_name\"`\n}"

	assert.Equal(t, actual, expected)
}
