package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"

	"github.com/naysayer/schemer/api/structure"
	"github.com/naysayer/schemer/app/postgres"
	"github.com/naysayer/schemer/app/postgres/table"
)

func main() {
	var structures []structure.Structure
	contents, err := ioutil.ReadFile("./schema.txt")
	if err != nil {
		log.Fatal(err)
	}

	regex := regexp.MustCompile("(?s)CREATE TABLE.*?\\);")
	located := regex.FindAll(contents, -1)
	for _, l := range located {
		tabLines := table.New(string(l))

		cluster, err := postgres.Cluster{}.New(tabLines)
		if err != nil {
			log.Println(err)
		}

		structures = append(structures, cluster)
	}
	PrintStrucutres(structures)
}

func PrintStrucutres(structures []structure.Structure) {
	for _, str := range structures {
		fmt.Print(structure.Stringify(str))
		fmt.Println()
		fmt.Println()
	}
}

func TableName(s string) string {
	r := regexp.MustCompile("(CREATE TABLE) | (\\()")
	return strings.Trim(r.ReplaceAllString(s, ""), " ")
}
