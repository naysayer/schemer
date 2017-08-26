package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"

	"github.com/naysayer/schemer/api"
	"github.com/naysayer/schemer/api/structure"
	"github.com/naysayer/schemer/app/postgres"
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
		tableLines := api.Seperate(string(l))

		cluster, err := postgres.Cluster{}.New(tableLines)
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
