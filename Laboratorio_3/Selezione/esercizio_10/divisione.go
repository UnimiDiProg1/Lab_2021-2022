package main

import "fmt"

func main() {
	var dividendo, divisore int

	var quoziente float64

	fmt.Println("Inserisci due numeri:")
	fmt.Scanln(&dividendo, &divisore)

	if divisore != 0 {
		quoziente = float64(dividendo) / float64(divisore)
		fmt.Println("Quoziente =", quoziente)
	} else {
		fmt.Println("Impossibile")
	}
}
