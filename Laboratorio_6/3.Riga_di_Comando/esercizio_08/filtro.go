package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	valori := LeggiNumeri()
	suff, insuff := FiltraVoti(valori)

	fmt.Println("Voti sufficienti:", suff)
	fmt.Println("Voti insufficienti:", insuff)

}

func LeggiNumeri() (numeri []int) {
	for _, v := range os.Args[1:] {
		if n, err := strconv.Atoi(v); err == nil {
			numeri = append(numeri, n)
		}
	}
	return
}

func FiltraVoti(valori []int) (sufficienti, insufficienti []int) {
	for _, v := range valori {
		if v >= 60 {
			sufficienti = append(sufficienti, v)
		} else {
			insufficienti = append(insufficienti, v)
		}
	}
	return
}
