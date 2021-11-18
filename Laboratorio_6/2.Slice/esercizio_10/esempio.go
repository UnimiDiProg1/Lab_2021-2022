package main

import (
	"fmt"
)

const Dimensione = 10

func main() {

	var a []int

	for i := 0; i < Dimensione; i++ {
		a = append([]int{i + 1}, a...)
	}

	fmt.Println(a)

	b := make([]int, Dimensione)

	copy(b, a[Dimensione/2:])

	fmt.Println(b)

}
