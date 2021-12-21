package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	n := 4
	var slc [][]int = CreaSliceBidimensionale(n)
	AssegnaSliceBidimensionale(slc)
	fmt.Println(slc)

	slc_soprasoglia := FiltraSliceBidimensionale(slc, 2)
	fmt.Println(slc_soprasoglia)
}

func CreaSliceBidimensionale(l int) [][]int {
	var s [][]int = make([][]int, l)
	for i := 0; i < l; i++ {
		s[i] = make([]int, l)
	}
	return s
}

func AssegnaSliceBidimensionale(s [][]int) {
	rand.Seed(int64(time.Now().Nanosecond()))
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s[i]); j++ {
			s[i][j] = rand.Intn(2)
		}
	}
}

func FiltraSliceBidimensionale(s [][]int, soglia int) [][]int {
	var risultato [][]int
	for _,riga := range s {
		somma := 0
		for _,valore := range riga {
			somma += valore
		}
		if somma >= soglia {
			risultato = append(risultato, riga)
		}
	}
	return risultato
}