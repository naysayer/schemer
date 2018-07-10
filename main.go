package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/naysayer/schemer/api/structure"
	"github.com/naysayer/schemer/app/postgres"
)

var (
	schema string
	sql    bool

	errEmptyFilepath = errors.New("please input a schema command line argument to the schema file you wish to parse")
)

func main() {
	flag.StringVar(&schema, "schema", "", "Location of directory that contains the desired config file.")
	flag.BoolVar(&sql, "sql", false, "parse either series of sql statements or a schema file, if true its sql")
	flag.Parse()
	if schema == "" {
		log.Fatal(errEmptyFilepath)
	}

	contents, err := ioutil.ReadFile(schema)
	if err != nil {
		log.Fatal(err)
	}

	structures, err := postgres.FromBytes(contents, sql)
	if err != nil {
		log.Fatal(err)
	}

	for _, str := range structures {
		fmt.Print(structure.Stringify(str), "\n\n")
	}
}
