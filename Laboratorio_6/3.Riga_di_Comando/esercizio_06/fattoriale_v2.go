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

func Fattoriali(n int) (listaFattoriali []int) {
	fattoriale := 1
	for i := 1; i <= n; i++ {
		fattoriale *= i
		listaFattoriali = append(listaFattoriali, fattoriale)
	}
	return
}
