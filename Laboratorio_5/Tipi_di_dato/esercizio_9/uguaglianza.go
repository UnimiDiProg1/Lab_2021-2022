/*
Testa la precisione del tipo di dato float

Effettua il confronto tra un numero ed il quadrato della sua radice quadrata per verificare la precisione del tipo float e la correttezza del risultato
*/
package main

import (
	"fmt"
	"math"
)

func main() {

	var x, y float64

	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&x)

	if y = math.Sqrt(x); x == y*y {
		fmt.Println(x, "uguale a", y, "*", y)
	} else {
		fmt.Println(x, "diverso da", y, "*", y)
	}

}
