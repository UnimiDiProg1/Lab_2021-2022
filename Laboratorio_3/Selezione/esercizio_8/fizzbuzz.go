package main

import "fmt"

func main() {

	var numero int

	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&numero)

	if numero%3 == 0 && numero%5 == 0 {
		// Se il numero è sia multiplo di 3 che di 5
		fmt.Print("Fizz Buzz")
	} else if numero%3 == 0 {
		// Se il numero è multiplo solo di 3 (si arriva qua SOLO SE l'espressione dell'if in riga 12 era falsa)
		fmt.Print("Fizz")
	} else if numero%5 == 0 {
		// Se il numero è multiplo solo di 5 (si arriva qua SOLO SE le espressioni degli if in riga 12 e 15 erano false)
		fmt.Print("Buzz")
	} else {
		// in questo caso il numero non è multiplo di 3 e non è multiplo di 5
		// questo ramo else è superfluo e può essere rimosso
	}

	fmt.Println()

}
