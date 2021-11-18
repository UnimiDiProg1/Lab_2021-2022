package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	var somma int

	numeri := LeggiNumeri()
	for i := range numeri {
		if Occorrenze(numeri, numeri[i]) == 1 {
			somma += numeri[i]
		}
	}

	fmt.Println(somma)
	/*
		var numeri [int]int
		numeri = make(map[int]int)
		for _, v := range os.Args[1:] {
			n, _ := strconv.Atoi(v)
			numeri[n]++
		}

		somma := 0
		for n, contatore := range numeri {
			if contatore == 1 {
				somma += n
			}
		}

		fmt.Println(somma)
	*/

}

func LeggiNumeri() (numeri []int) {
	for _, v := range os.Args[1:] {
		if n, err := strconv.Atoi(v); err == nil {
			numeri = append(numeri, n)
		}
	}
	return
}

func Occorrenze(numeri []int, n int) int {
	contatore := 0
	for i := range numeri {
		if numeri[i] == n {
			contatore++
		}
	}
	return contatore
}
