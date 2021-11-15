package main

import "fmt"

func main() {

	var a, b int
	fmt.Scan(&a, &b)

	var c ???

	c = a == b

	if c {
		fmt.Println("uguali")
	} else {
		fmt.Println("diversi")
	}

}
