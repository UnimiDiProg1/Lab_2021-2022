package main

import "fmt"

func main() {
	var angolo1, angolo2 float64

	fmt.Print("Inserire le ampiezze dei due angoli: ")
	fmt.Scan(&angolo1, &angolo2)

	var sommaAmpiezze float64
	sommaAmpiezze = angolo1 + angolo2

	if sommaAmpiezze < 180 {
		fmt.Print("Ampiezza terzo angolo = ",
			180-sommaAmpiezze, "°\n")
	} else {
		fmt.Println("I due angoli non appartengono ad un triangolo")
	}
}
