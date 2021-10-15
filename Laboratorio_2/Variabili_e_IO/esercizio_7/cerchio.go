package main

import "fmt"
import "math"

func main() {

	var raggio float64
	var circonferenza, area float64

	fmt.Print("Raggio = ")
	fmt.Scan(&raggio)

	circonferenza = raggio * 2 * math.Pi
	area = raggio * raggio * math.Pi

	fmt.Println("Circonferenza =", circonferenza)
	fmt.Println("Area =", area)

}
