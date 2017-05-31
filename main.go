package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"

	"github.com/naysayer/schemer/app"
	"github.com/naysayer/schemer/table"
)

func main() {
	contents, err := ioutil.ReadFile("./schema.txt")
	if err != nil {
		log.Fatal(err)
	}

	regex := regexp.MustCompile("(?s)CREATE TABLE.*?\\);")
	located := regex.FindAll(contents, -1)
	for _, l := range located {
		tabLines := table.New(string(l))

		cluster, err := app.Cluster{}.New(tabLines)
		if err != nil {
			log.Println(err)
		}

		fmt.Println(cluster.Table)
		for _, c := range cluster.Columns {
			fmt.Println(c.Name, c.Type)
		}
	}
}

func TableName(s string) string {
	r := regexp.MustCompile("(CREATE TABLE) | (\\()")
	return strings.Trim(r.ReplaceAllString(s, ""), " ")
}
