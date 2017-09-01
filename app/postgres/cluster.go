// A cluster resolves the structure interface. It is used to represent
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
	patternCreateTable = regexp.MustCompile("(?s)CREATE TABLE.*?\\);")
)

type cluster struct {
	Table   string
	Columns []*column.Column
}

// Title returns the title name of the table as the title. This is capitalized
// as these are used to represent the names of structs, so we want to export
// them by default.
func (c *cluster) Title() string {
	return strings.Title(c.Table)
}

func (c *cluster) Contents() []attr.Attr {
	var atts []attr.Attr
	for _, v := range c.Columns {
		atts = append(atts, v)
	}
	return atts
}

func New(lines []string) (structure.Structure, error) {
	var columns []*column.Column

	for _, l := range lines {
		col, err := column.Column{}.Populate(l)
		if err != nil && err != column.EndingOfCreateError && err != column.BeginningOfCreateError {
			return nil, err
		}

		if err != column.EndingOfCreateError && err != column.BeginningOfCreateError {
			columns = append(columns, col)
		}
	}

	return &cluster{Table: table.Name(lines), Columns: columns}, nil
}

func NewFromBytes(bytes []byte) ([]structure.Structure, error) {
	var structures []structure.Structure

	for _, l := range patternCreateTable.FindAll(bytes, -1) {
		tableLines := api.Seperate(string(l))

		cluster, err := New(tableLines)
		if err != nil {
			return structures, err
		}

		structures = append(structures, cluster)
	}

	return structures, nil
}
