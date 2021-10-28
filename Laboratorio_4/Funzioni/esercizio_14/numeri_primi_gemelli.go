package main

import (
	"fmt"
)

func main() {
	var soglia int
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&soglia)

	if soglia > 0 {
		NumeriPrimiGemelli(soglia)
	} else {
		fmt.Println("La soglia inserita non è positiva")
	}

}

func ÈPrimo(n int) bool {
	var divisore int
	for divisore = 2; divisore < n; divisore++ {
		if n%divisore == 0 {
			return false
		}
	}
	return true
}

func NumeriPrimiGemelli(n int) {
	fmt.Println("Numeri primi gemelli inferiori a", n)
	var numero int
	for numero = 2; numero < n-2; numero++ {
		if ÈPrimo(numero) && ÈPrimo(numero+2) {
			fmt.Print("(", numero, ",", numero+2, ") ")
		}
	}
	fmt.Println()
}
