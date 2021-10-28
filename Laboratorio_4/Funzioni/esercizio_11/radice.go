package main

import (
	"fmt"
	"math"
)

func main() {

	var n float64
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&n)

	var radice float64
	var ok bool

	radice, ok = RadiceQuadrata(n)
	if ok {
		fmt.Println("Radice quadrata:", radice)
	} else {
		fmt.Println("Il valore inserito non appartiene al dominio della funzione")
	}

}

func RadiceQuadrata(numero float64) (float64, bool) {
	if numero >= 0 {
		return math.Sqrt(numero), true
	} else {
		return 0, false
	}
}
