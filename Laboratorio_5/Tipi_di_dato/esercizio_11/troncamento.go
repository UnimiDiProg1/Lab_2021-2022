package main

import (
	"fmt"
	"math"
)

func main() {

	var valore float64

	fmt.Print("Inserisci il valore da troncare: ")
	fmt.Scan(&valore)

	valoreTroncato := math.Trunc(valore*100) / 100
	fmt.Println("Valore troncato =", valoreTroncato)

}
