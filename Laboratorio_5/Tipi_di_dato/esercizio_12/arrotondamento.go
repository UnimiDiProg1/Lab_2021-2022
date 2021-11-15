package main

import (
	"fmt"
	"math"
)

func main() {

	var valore float64

	fmt.Print("Inserisci il valore da arrotondare: ")
	fmt.Scan(&valore)

	valoreArrotondato := math.Round(valore*100) / 100
	fmt.Println("Valore arrotondato =", valoreArrotondato)

}
