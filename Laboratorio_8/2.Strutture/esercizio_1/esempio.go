package main

import (
	"fmt"
)

type Persona struct {
	Nome    string
	Cognome string
}

func main() {

	var p1 Persona
	p1 = Persona{"Rick", "Sanchez"}

	var p2 Persona
	p2.Nome = "Jerry"
	p2.Cognome = "Smith"

	p3 := Persona{Nome: "Morty"}

	p2 = p3

	fmt.Println(p1, p2)
}
