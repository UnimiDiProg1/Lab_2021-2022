package main

import "fmt"

func main() {

	var distanzaKm float64
	var distanzaMiglia float64

	fmt.Print("Distanza (Km) = ")
	fmt.Scanln(&distanzaKm)

	distanzaMiglia = distanzaKm * 0.62137

	fmt.Println("Distanza (mi) =", distanzaMiglia)

}
