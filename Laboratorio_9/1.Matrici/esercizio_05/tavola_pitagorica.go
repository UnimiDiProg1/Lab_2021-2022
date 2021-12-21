package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	tavola := CreaTavolaPitagorica(LeggiNumero())
	StampaTavolaPitagorica(tavola)

}

func LeggiNumero() (n int) {
	n, _ = strconv.Atoi(os.Args[1])
	return
}

func CreaTavolaPitagorica(n int) (tavolaPitagorica [][]int) {
	tavolaPitagorica = make([][]int, n)
	for i := range tavolaPitagorica {
		tavolaPitagorica[i] = make([]int, n)
		for j := range tavolaPitagorica[i] {
			tavolaPitagorica[i][j] = (i + 1) * (j + 1)
		}
	}
	return
}

func StampaTavolaPitagorica(tavolaPitagorica [][]int) {
	for _, riga := range tavolaPitagorica {
		for _, valore := range riga {
			fmt.Printf("%4d ", valore)
		}
		fmt.Println()
	}
	return
}
