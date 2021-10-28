package main

import "fmt"

func main() {

	var soglia int
	fmt.Print("Inserisci la soglia: ")
	fmt.Scan(&soglia)

	TernePitagoriche(soglia)

}

func ÈTernaPitagorica(a, b, c int) bool {
	return a*a+b*b == c*c
}

func TernePitagoriche(soglia int) {
	var a, b, c int

	fmt.Println("Terne pitagoriche:")
	for a = 1; a < soglia; a++ {
		for b = 1; b < soglia; b++ {
			for c = 1; c < soglia; c++ {
				if ÈTernaPitagorica(a, b, c) {
					fmt.Print("(", a, ", ", b, ", ", c, ")\n")
				}
			}
		}
	}
}
