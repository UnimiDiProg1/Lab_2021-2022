package main

import "fmt"

func main() {
	var a [6]int

	for i := range a {
		a[i] = i
	}

	var b []int
	b = a[:]

	for i := range b {
		b[i] = i * 2
	}

	fmt.Println(a)
}
