package main

import "fmt"

func main() {

	var soglia int
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&soglia)

	if soglia > 0 {
		NumeriAbbondanti(soglia)
	} else {
		fmt.Println("La soglia inserita non è positiva")
	}

}

func SommaDivisoriPropri(n int) (somma int) {
	var divisore int
	for divisore = 1; divisore < n; divisore++ {
		if n%divisore == 0 {
			somma += divisore
		}
	}
	return
}

func ÈAbbondante(n int) bool {
	return n > 0 && SommaDivisoriPropri(n) > n
}

func NumeriAbbondanti(soglia int) {
	fmt.Print("Numeri abbondanti:")
	var numero int
	for numero = 1; numero < soglia; numero++ {
		if ÈAbbondante(numero) {
			fmt.Print(" ", numero)
		}
	}
	fmt.Println()
}
