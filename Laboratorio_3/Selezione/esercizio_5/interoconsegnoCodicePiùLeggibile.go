package main

import "fmt"

func main() {

	var numero int

	fmt.Print("Inserisci numero: ")
	fmt.Scan(&numero)

	if numero > 0 {
		fmt.Print("+")
	}

	fmt.Println(numero)

}
