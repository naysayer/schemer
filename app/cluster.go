package app

import (
	"github.com/naysayer/schemer/table"
	"github.com/naysayer/schemer/table/column"
)

type Cluster struct {
	Table   string
	Columns []*column.Column
}

// func (c *Cluster) Print() {
// 	fmt.Printf("Type %s struct {\n", c.Table)
// 	for _, col := range c.Columns {
// 		fmt.Printf("Type %s struct {\n", col.Name)
// 		fmt.Printf("\t ")
// 	}
// 	fmt.Print("}")
// }

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
