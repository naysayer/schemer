package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"

	"github.com/naysayer/schemer/api/table"
	"github.com/naysayer/schemer/app"
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
		for _, c := range cluster.Contents() {
			fmt.Println(c.Title(), c.Type(), c.Tags())
		}
		fmt.Println()
	}
}

func TableName(s string) string {
	r := regexp.MustCompile("(CREATE TABLE) | (\\()")
	return strings.Trim(r.ReplaceAllString(s, ""), " ")
}
