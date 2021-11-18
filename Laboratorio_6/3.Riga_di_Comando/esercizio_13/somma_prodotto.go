package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	numeri := LeggiNumeri()
	fmt.Println("La somma è", Calcola(numeri))
}

func LeggiNumeri() (numeri []int) {
	for _, v := range os.Args[1:] {
		if n, err := strconv.Atoi(v); err == nil {
			numeri = append(numeri, n)
		}
	}
	return
}

func Calcola(numeri []int) (somma int) {
	for i := 0; i < len(numeri); i += 2 {
		somma += numeri[i] * numeri[i+1]
	}
	return
}
