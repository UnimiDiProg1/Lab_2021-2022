/*
Stampa un quadrato di * e +

Dato un numero n letto a tastiera, stampa un quadrato composto da n righe di * e + alternate
*/
package main

import (
	"fmt"
)

func main() {

	var numero int
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&numero)

	var i, j int
	for i = 0; i < numero; i++ {
		for j = 0; j < numero; j++ {

			// Possiamo distinguere i due casi * e + semplicemente controllando se l'indice di riga
			// è multiplo di 2. In questo modo possiamo alternare delle sequenze prestabilite
			// In questo caso, con righe di indice pari stamperemo *, mentre con righe di indice dispari dei +
			if i%2 == 0 {
				fmt.Print("* ")
			} else {
				fmt.Print("+ ")
			}
		}
		fmt.Println()
	}

}
