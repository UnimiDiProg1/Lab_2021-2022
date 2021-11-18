package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	valori := LeggiNumeri()
	fmt.Println("La somma è", Calcola(valori))
}

func LeggiNumeri() (numeri []int) {
	for _, v := range os.Args[1:] {
		if n, err := strconv.Atoi(v); err == nil {
			numeri = append(numeri, n)
		}
	}
	return
}

func Calcola(valori []int) (somma int) {
	for i, v1 := range valori {
		for _, v2 := range valori[i+1:] {
			if (v1*v2)%2 == 0 {
				somma += v1 * v2
			}
		}
	}
	return
}
