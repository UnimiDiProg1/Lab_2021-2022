package main

import "fmt"

func main() {
	var base, altezza float64
	var perimetro, area float64

	// lettura input
	fmt.Print("Inserisci la base: ")
	fmt.Scan(&base)

	fmt.Print("Inserisci l'altezza: ")
	fmt.Scan(&altezza)

	// calcolo perimetro e area
	perimetro = 2 * (base + altezza)
	area = base * altezza

	// stampa dell'output
	fmt.Println("Perimetro =", perimetro)
	fmt.Println("Area =", area)
}
