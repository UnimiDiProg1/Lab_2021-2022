package main

import (
	"fmt"
)

func main() {

	var numero int
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&numero)

	// stampa n righe di n caratteri * e + alternati
	var i, j int
	for i = 0; i < numero; i++ {
		for j = 0; j < numero; j++ {

			// Questo codice è identico al codice di quadratiAlternati.go:
			// in quel codice si distingueva tra righe di indice pari e di indice dispari.
			// In questo caso però, si devono alternare colonne diverse e non righe, quindi l'indice da usare
			// è quello di colonna
			if j%2 == 0 {
				fmt.Print("* ")
			} else {
				fmt.Print("+ ")
			}
		}
		fmt.Println()
	}

}
