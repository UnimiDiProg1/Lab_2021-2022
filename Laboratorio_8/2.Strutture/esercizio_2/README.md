# Qual è l'output?

Qual è l'output del seguente programma?

```go
package main

import (
	"fmt"
)

type Persona struct {
	Nome    string
	Cognome string
}

func main() {

	anagrafica := []Persona{
		{"Rick", "Sanchez"},
		{"Morty", "Smith"},
		{"Jerry", "Smith"},
		{"Summer", "Smith"},
		{"Beth", "Smith"}}

	CambiaPrimoCognome(anagrafica)

	CambiaCognome(anagrafica[0])

	fmt.Println(anagrafica[0])
}

func CambiaPrimoCognome(a []Persona) {
	a[0].Cognome = "Martinez"
}

func CambiaCognome(p Persona) {
	p.Cognome = "Gomez"
}
```
