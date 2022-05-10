package main

import (
	"fmt"

	"github.com/arxxm/API_catfact.git/catfacts"
)

func main() {
	client := catfacts.NewClient()
	fact, err := client.ListAllFacts()
	if err != nil {
		panic(err)
	}
	for _, onefact := range fact {
		fmt.Println(onefact.Fact)
	}
	// fmt.Println(fact.Fact)
}
