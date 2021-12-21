package main

import (
	"example.com/main/triangolo"
	"fmt"
	"os"
	"strconv"
)

func main() {

	lati := [3]float64{}
	for i, v := range os.Args[1:] {
		lati[i], _ = strconv.ParseFloat(v, 64)
	}

	t, ok := triangolo.NuovoTriangolo(lati[0], lati[1], lati[2])

	if !ok {
		fmt.Println("Triangolo non valido")
		return
	} else {
		fmt.Println("Perimetro triangolo =", triangolo.Perimetro(t))
		fmt.Println("Area triangolo =", triangolo.Area(t))
	}	
}
