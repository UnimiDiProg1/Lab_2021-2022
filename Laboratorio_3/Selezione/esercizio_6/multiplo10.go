package main

import "fmt"

func main() {

	var numero int

	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&numero)

	fmt.Print(numero)

	if numero%10 != 0 {
		fmt.Print(" non")
	}

	fmt.Println(" è multiplo di 10")
}
