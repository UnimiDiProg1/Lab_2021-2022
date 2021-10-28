/*
Stampa un quadrato di asterischi

Dato un numero n letto a tastiera, stampa un quadrato n x n di * intervallati da spazi
*/
package main

import "fmt"

func main() {

	var numero int

	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&numero)

	// Il primo for serve per scorrere le righe del quadrato
	var i, j int
	for i = 0; i < numero; i++ {
		// Il secondo for per le colonne del quadrato
		for j = 0; j < numero; j++ {
			fmt.Print("* ")
		}

		// Al termine di ogni riga è necessario stampare un 'a capo'
		// altrimenti tutti gli asterischi sarebbero stampati su un'unica riga
		fmt.Println()
	}

}
