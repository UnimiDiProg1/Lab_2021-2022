/*
Stampa un quadrato di asterischi

Dato un numero n letto a tastiera, stampa un quadrato n x n con * sui bordi e + al centro
*/
package main

import "fmt"

func main() {

	var numero int

	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&numero)

	var i, j int
	for i = 0; i < numero; i++ {
		for j = 0; j < numero; j++ {

			// Il programma deve stampare un asterisco solo sui bordi del quadrato.
			// Per il bordo superiore e quello di sinistra basta controllare semplicemente
			// se gli indici hanno valore uguale a 0 (la prima riga o la prima colonna)
			// Per il bordo inferiore e quello di destra, in modo analogo si controlla
			// che l'indice sia uguale alla dimensione del quadrato -1
			if i == 0 || j == 0 || i == numero-1 || j == numero-1 {
				fmt.Print("* ")
			} else {
				fmt.Print("+ ")
			}
		}
		fmt.Println()
	}

}
