package main

import "fmt"

func funzione(x int) int {
	return x + 5
}

func main() {
	var x int = 10
	fmt.Println(funzione(x))
}
