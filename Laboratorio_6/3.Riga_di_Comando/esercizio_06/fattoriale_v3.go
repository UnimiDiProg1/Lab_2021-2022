package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {

	n := LeggiNumero()

	start := time.Now()
	fattoriali1 := FattorialiV1(n)
	end := time.Now()
	fmt.Printf("Fattoriali V1: %v - Tempo %d\n", fattoriali1, end.Sub(start))

	start = time.Now()
	fattoriali2 := FattorialiV2(n)
	end = time.Now()
	fmt.Printf("Fattoriali V2: %v - Tempo %d\n", fattoriali2, end.Sub(start))
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

func FattorialiV1(n int) []int {
	listaFattoriali := make([]int, n)
	for i := 1; i <= n; i++ {
		listaFattoriali[i-1] = Fattoriale(i)
	}
	return listaFattoriali
}

func FattorialiV2(n int) (listaFattoriali []int) {
	fattoriale := 1
	listaFattoriali = make([]int, n)
	for i := 1; i <= n; i++ {
		fattoriale *= i
		listaFattoriali[i-1] = fattoriale
	}
	return
}
