package main

import "fmt"

func main() {
	var numero float64 = 0
	fmt.Println(numero)

	fmt.Println(0 / numero)

	numero = 1 / numero
	fmt.Println(numero)

	numero = 1 / numero
	fmt.Println(numero)

	numero = -(1 / numero)
	fmt.Println(numero)

	fmt.Println(1 / numero)

	numero = numero - numero
	fmt.Println(numero)
}

