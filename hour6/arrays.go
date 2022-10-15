package main

import (
	"fmt"
)

func main() {
	var cheeses = make([]string, 2)
	cheeses[0] = "Mariolles"
	cheeses[1] = "Epoisses de Bourgogne"
	// cheeses = append(cheeses, "Camembert", "oui", "non")
	// fmt.Println(len(cheeses))
	// fmt.Println(cheeses)
	// cheeses = append(cheeses[:2], cheeses[2+1:]...)
	// fmt.Println(len(cheeses))
	// fmt.Println(cheeses)

	var smellyCheeses = make([]string, 2)
	copy(smellyCheeses, cheeses)
	fmt.Println(smellyCheeses)

	var players = make(map[string]int)
	players["cook"] = 32
	players["bairstow"] = 27
	players["stokes"] = 26
	delete(players, "cook")
	fmt.Println(players)

	var HTMLElements = make(map[string]string)
	HTMLElements["p"] = "Paragraph"
	HTMLElements["img"] = "Image"
	HTMLElements["h1"] = "Heading One"
	HTMLElements["h2"] = "Heading Two"
	fmt.Println(HTMLElements)

}
