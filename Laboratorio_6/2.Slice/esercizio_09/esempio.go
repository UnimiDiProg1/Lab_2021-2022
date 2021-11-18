package main

import (
	"fmt"
)

func main() {

	s1 := make([]int, 4, 8)

	for i := 0; i < len(s1); i++ {
		s1[i] = i + 1
	}

	fmt.Println("1)\ns1:")
	stampa(s1)

	s2 := append(s1[2:], []int{10, 20}...)
	fmt.Println("\n2)\ns2:")
	stampa(s2)
	fmt.Println("s1:")
	stampa(s1)

	s2[0] = -1
	s2 = append(s2, 100)
	fmt.Println("\n3)\ns2:")
	stampa(s2)
	fmt.Println("s1:")
	stampa(s1)

	s2 = append(s2, 1000, 2000)
	fmt.Println("\n4)\ns2:")
	stampa(s2)
	fmt.Println("s1:")
	stampa(s1)

	s2 = append(s2[:2], s2[len(s2)-2:]...)
	fmt.Println("\n5)\ns2:")
	stampa(s2)

}

func stampa(s []int) {
	var lunghezza int
	fmt.Println("a) ", s, len(s), cap(s))
	lunghezza = len(s)
	s = s[:cap(s)]
	fmt.Println("b) ", s, len(s), cap(s))
	s = s[:lunghezza]
}
