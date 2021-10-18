/*
Operazioni numeriche

Effettua delle operazioni tra due numeri interi inseriti da tastiera.
Le operazioni sono:
- individuazione del minimo e del massimo
- somma, sottrazione, prodotto e divisione
- elevamento a potenza
- media
*/
package main

import (
	"fmt"
	"math"
)

func main() {

	var x, y int

	fmt.Println("Inserisci due numeri interi: ")
	fmt.Scan(&x, &y)

	var max int = x
	if max < y {
		max = y
	}

	var min int = x
	if min > y {
		min = y
	}

	fmt.Println("Maggiore:", max)
	fmt.Println("Minore:", min)

	fmt.Println("Somma:", x+y)
	fmt.Println("Differenza:", max-min)
	fmt.Println("Prodotto:", x*y)
	fmt.Println("Divisione:", float64(x)/float64(y))

	fmt.Println("Valore medio:", (x+y)/2)

	var potenza int = 1
	
	var i int;
	for i = 0; i < y; i++ {
		potenza *= x
	}

	fmt.Println("Potenza (ciclo for):", potenza)
	fmt.Println("Potenza (math.Pow):", int(math.Pow(float64(x), float64(y))))
}
