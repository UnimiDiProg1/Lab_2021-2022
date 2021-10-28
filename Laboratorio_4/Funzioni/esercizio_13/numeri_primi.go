// Stampa la sequenza di numeri primi inferiori ad una soglia data da standard input.
package main

import (
	"fmt"
)

func main() {
	var soglia int
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&soglia)

	if soglia > 0 {
		NumeriPrimi(soglia)
	} else {
		fmt.Println("La soglia inserita non è positiva")
	}

}

// ÈPrimo calcola se un numero passato come argomento è primo.
// Restituisce true in caso affermativo, false altrimenti.
func ÈPrimo(n int) bool {
	var divisore int
	for divisore = 2; divisore < n; divisore++ {
		if n%divisore == 0 {
			return false
		}
	}
	return true
}

// NumeriPrimi stampa a standard output la sequenza di numeri interi primi inferiori alla soglia passata come argomento.
func NumeriPrimi(n int) {
	fmt.Println("Numeri primi inferiori a", n)
	var numero int
	for numero = 2; numero < n; numero++ {
		if ÈPrimo(numero) {
			fmt.Print(numero, " ")
		}
	}
	fmt.Println()
}
