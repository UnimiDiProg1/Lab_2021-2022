package main

import "fmt"
import "math"

func main() {

	var raggio float64
	fmt.Print("Inserisci il raggio del cerchio: ")
	fmt.Scan(&raggio)

	fmt.Println("Area del cerchio:", CalcolaArea(raggio))
	fmt.Println("Circonferenza del cerchio:", CalcolaCirconferenza(raggio))

}

func CalcolaArea(raggio float64) float64 {
	return math.Pi * raggio * raggio
}

func CalcolaCirconferenza(raggio float64) float64 {
	return 2 * raggio * math.Pi
}
