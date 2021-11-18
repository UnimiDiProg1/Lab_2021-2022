package main

import "fmt"

func main() {
	var a []int
	a = []int{0, 1, 2, 3, 4, 5, 6}

	var b []int
	b = a[2:4]

	b[0] = a[0]
	b[len(b)-1] = a[0]

	fmt.Println(a)
}
