package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	n := LeggiNumero()
	fmt.Println("Fattoriali:", Fattoriali(n))

}

func LeggiNumero() (n int) {
	n, _ = strconv.Atoi(os.Args[1])
	return
}

func Fattoriale(n int) int {
	fattoriale := 1
	for i := 1; i <= n; i++ {
		fattoriale *= i
	}
	return fattoriale
}

func Fattoriali(n int) []int {
	listaFattoriali := make([]int, n)
	for i := 1; i <= n; i++ {
		listaFattoriali[i-1] = Fattoriale(i)
	}
	return listaFattoriali
}
