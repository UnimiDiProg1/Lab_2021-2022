/*
Calcola la tabellina di un numero

Letto un numero n da tastiera, il package calcola la sua tabellina
*/
package main

import "fmt"

func main() {
	var numero int

	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&numero)

	var i int;
	for i = 1; i <= 10; i++ {
		fmt.Println(i, "x", numero, "=", i*numero)
	}
}
