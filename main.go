package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/naysayer/schemer/api/structure"
	"github.com/naysayer/schemer/app/postgres"
)

func main() {
	contents, err := ioutil.ReadFile("./schema.txt")
	if err != nil {
		log.Fatal(err)
	}

	structures, err := postgres.FromBytes(contents)
	if err != nil {
		log.Fatal(err)
	}
	printStrucutres(structures)
}

func printStrucutres(structures []structure.Structure) {
	for _, str := range structures {
		fmt.Print(structure.Stringify(str))
		fmt.Println()
		fmt.Println()
	}
}
