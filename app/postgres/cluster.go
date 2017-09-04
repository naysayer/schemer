// A postgres resolves the structure interface. It is used to represent
// a postgres sql database schema file. Its internals build on top of table
// and column, thus the strucutre of that directory and the corresponding naming
// needs to be changed.

package postgres

import (
	"regexp"
	"strings"

	"github.com/naysayer/schemer/api"
	"github.com/naysayer/schemer/api/structure"
	"github.com/naysayer/schemer/api/structure/attr"
	"github.com/naysayer/schemer/app/postgres/table"
	"github.com/naysayer/schemer/app/postgres/table/column"
)

var (
	patternCreateTable = regexp.MustCompile(`(?s)CREATE TABLE.*?\);`)
)

// postgres resolves the structure interface. It is used to represent
// a postgres sql database schema file. Its internals build on top of table
// and column. The name represents the table's name, and the columns represent
// the subsiquent columns within said table.
type postgres struct {
	Table   string
	Columns []*column.Column
}

// Title returns the title name of the table as the title. This is capitalized
// as these are used to represent the names of structs, so we want to export
// them by default.
func (c *postgres) Title() string {
	return strings.Title(c.Table)
}

func (c *postgres) Contents() []attr.Attr {
	var atts []attr.Attr
	for _, v := range c.Columns {
		atts = append(atts, v)
	}
	return atts
}

// New returns a struct of postgres that conforms to the structure.Structure
// interface. It does this by taking a slice of strings that represents
// a create table statement within a postgres database schema
func New(lines []string) (structure.Structure, error) {
	var columns []*column.Column

	for _, l := range lines {
		col, err := column.New(l)
		if (err != nil) && (err != column.ErrEndingOfCreate && err != column.ErrBeginningOfCreate) {
			return nil, err
		}

		if err != column.ErrEndingOfCreate && err != column.ErrBeginningOfCreate {
			columns = append(columns, col)
		}
	}

	return &postgres{Table: table.Name(lines[0]), Columns: columns}, nil
}

// FromBytes given a slice of bytes returns a structure.Structure interface
// it does this by leveraging the New function accordingly.
func FromBytes(bytes []byte) ([]structure.Structure, error) {
	var structures []structure.Structure

	for _, l := range patternCreateTable.FindAll(bytes, -1) {
		tableLines := api.Seperate(string(l))

		postgres, err := New(tableLines)
		if err != nil {
			return nil, err
		}

		structures = append(structures, postgres)
	}

	return structures, nil
}
