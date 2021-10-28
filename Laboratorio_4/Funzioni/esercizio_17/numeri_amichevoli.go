package main

import "fmt"

func main() {

	var soglia int
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&soglia)

	if soglia > 0 {
		NumeriAmichevoli(soglia)
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

func SonoAmichevoli(n, m int) bool {
	return SommaDivisoriPropri(n) == m && SommaDivisoriPropri(m) == n
}

func NumeriAmichevoli(soglia int) {
	fmt.Println("Numeri amichevoli inferiori a", soglia)
	var numeroX, numeroY int
	for numeroX = 1; numeroX < soglia; numeroX++ {
		for numeroY = numeroX + 1; numeroY < soglia; numeroY++ {
			if SonoAmichevoli(numeroX, numeroY) {
				fmt.Print("(", numeroX, ",", numeroY, ") ")
			}
		}
	}
	fmt.Println()
}
