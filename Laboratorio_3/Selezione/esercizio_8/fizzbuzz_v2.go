package main

import "fmt"

func main() {

	var numero int

	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&numero)

	if numero%3 == 0 {
		fmt.Print("Fizz ")
	}
	if numero%5 == 0 {
		fmt.Print("Buzz ")
	}
	// Questa versione funziona solamente perché non si stampa nulla quando il numero non è multiplo né di 3 né di 5
	fmt.Println()

}
