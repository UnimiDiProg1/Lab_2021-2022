package main

import "fmt"

func main() {
	var numero int

	fmt.Print("Inserisci numero: ")
	fmt.Scan(&numero)

	fmt.Print("Divisori di ", numero, ": ")

	var i int;
	for i = 1; i <= numero/2; i++ {
		if numero % i == 0 {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}
