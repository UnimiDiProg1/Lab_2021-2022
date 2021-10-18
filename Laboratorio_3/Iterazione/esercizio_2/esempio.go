package main

import "fmt"

func main() {

	var n int

	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&n)

	var i int;
	for i = 1; i <= n; i *= 2 {
		fmt.Print(i, " ")
	}

	fmt.Println()

}
