package main

import "fmt"

func main() {

	const larghezzaTronco = 3
	const altezzaTronco = 3

	var n int
	fmt.Scan(&n)

	var larghezzaAlbero int = 2*n - 1

	var offsetTronco int = (larghezzaAlbero - larghezzaTronco) / 2
	if offsetTronco < 0 {
		offsetTronco = 0
	}

	var offsetAlbero int = (larghezzaTronco - larghezzaAlbero) / 2
	if offsetAlbero < 0 {
		offsetAlbero = 0
	}

	// Stampa albero
	var i, j int
	for i = 0; i < n; i++ {
		for j = 0; j < n-1-i+offsetAlbero; j++ {
			fmt.Print(" ")
		}
		for j = 0; j < 2*i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	// Stampa tronco
	for i = 0; i < altezzaTronco; i++ {
		for j = 0; j < offsetTronco; j++ {
			fmt.Print(" ")
		}
		for j = 0; j < larghezzaTronco; j++ {
			fmt.Print("+")
		}
		fmt.Println()
	}

}
