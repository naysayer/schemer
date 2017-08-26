// A cluster resolves the structure interface. It is used to represent
// a postgres sql database schema file
// Its internals build on top of table and column, thus the strucutre of that
// directory and the corresponding naming needs to be changed.

package postgres

import (
	"strings"

	"github.com/naysayer/schemer/api/structure/attr"
	"github.com/naysayer/schemer/app/postgres/table"
	"github.com/naysayer/schemer/app/postgres/table/column"
)

type Cluster struct {
	Table   string
	Columns []*column.Column
}

func (c Cluster) New(lines []string) (*Cluster, error) {
	var columns []*column.Column
	table := table.Name(lines)

	for _, l := range lines {
		col, err := column.Column{}.Populate(l)
		if err != nil && err != column.EndingOfCreateError && err != column.BeginningOfCreateError {
			return &Cluster{}, err
		}

		if err != column.EndingOfCreateError && err != column.BeginningOfCreateError {
			columns = append(columns, col)
		}
	}

	return &Cluster{Table: table, Columns: columns}, nil
}

// Title returns the title name of the table as the title. This is capitalized
// as these are used to represent the names of structs, so we want to export
// them by default.
func (c *Cluster) Title() string {
	return strings.Title(c.Table)
}

func (c *Cluster) Contents() []attr.Attr {
	var atts []attr.Attr
	for _, v := range c.Columns {
		atts = append(atts, v)
	}
	return atts
}
