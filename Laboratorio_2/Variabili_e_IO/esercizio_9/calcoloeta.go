package main

import "fmt"

func main() {
	var età1, età2 int

	fmt.Print("Età persona 1 = ")
	fmt.Scanln(&età1)
	fmt.Print("Età persona 2 = ")
	fmt.Scanln(&età2)

	var somma int
	var media float64

	somma = età1 + età2
	media = float64(età1+età2) / 2

	fmt.Println("Somma età =", somma)
	fmt.Println("Media età =", media)

	età1 += 10
	età2 += 10

	var somma10anni int
	var media10anni float64

	somma10anni = età1 + età2
	media10anni = float64(età1+età2) / 2

	fmt.Println("Somma età a 10 anni =", somma10anni)
	fmt.Println("Media età a 10 anni =", media10anni)
}
