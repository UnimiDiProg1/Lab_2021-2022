package main

import "fmt"

func main() {

	var a, b int

	fmt.Scan(&a, &b)

	var somma int

	var i int;
	for i = a + 1; i < b; i++ {
		if i%2 == 0 {
			somma += i
		}
	}

	fmt.Println("Somma =", somma)
}
