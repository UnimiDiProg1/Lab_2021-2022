package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {

	numeri := LeggiNumeri()
	fmt.Println("Minimo:", Minimo(numeri))
	fmt.Println("Massimo:", Massimo(numeri))
	fmt.Printf("Media: %.2f", Media(numeri))

}

func LeggiNumeri() (numeri []int) {
	numeri = make([]int, len(os.Args)-1)
	for i := 1; i < len(os.Args); i++ {
		numeri[i-1], _ = strconv.Atoi(os.Args[i])
	}
	return
}

func Minimo(numeri []int) int {
	min := math.MaxInt64
	for _, n := range numeri {
		if min > n {
			min = n
		}
	}
	return min
}

func Massimo(numeri []int) int {
	max := math.MinInt64
	for _, n := range numeri {
		if max < n {
			max = n
		}
	}
	return max
}

func Media(numeri []int) (media float64) {
	for _, n := range numeri {
		media += float64(n)
	}
	media /= float64(len(numeri))
	return
}
