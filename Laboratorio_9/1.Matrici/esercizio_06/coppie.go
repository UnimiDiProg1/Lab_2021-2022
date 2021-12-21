package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {

	sliceBidimensionale := CreaSliceBidimensionale(LeggiNumero())
	AssegnaSliceBidimensionale(sliceBidimensionale)
	fmt.Println("la slice originale è:")
	StampaSliceBidimensionale(sliceBidimensionale)

	fmt.Println("gli indirizzi degli elementi uguali a 1 sono:")
	coppie := Coppie(sliceBidimensionale)
	for _, coppia := range coppie {
		fmt.Println(coppia)
	}
}

func LeggiNumero() (n int) {
	n, _ = strconv.Atoi(os.Args[1])
	return
}

func CreaSliceBidimensionale(dimensione int) (risultato [][]int) {
	risultato = make([][]int, dimensione)
	for i := 0; i < dimensione; i++ {
		risultato[i] = make([]int, dimensione)
	}
	return
}

func AssegnaSliceBidimensionale(sliceBidimensionale [][]int) {
	rand.Seed(int64(time.Now().Nanosecond()))
	for i := range sliceBidimensionale {
		for j := range sliceBidimensionale[i] {
			sliceBidimensionale[i][j] = rand.Intn(2)
		}
	}
}

func StampaSliceBidimensionale(sliceBidimensionale [][]int) {
	for _, riga := range sliceBidimensionale {
		for _, valore := range riga {
			fmt.Print("\t", valore)
		}
		fmt.Println()
	}
}

func Coppie(sliceBidimensionale [][]int) (coppie [][]int) {
	for i, riga := range sliceBidimensionale {
		for j, valore := range riga {
			if valore == 1 {
				coppie = append(coppie, []int{i, j})
			}
		}
	}
	return
}
