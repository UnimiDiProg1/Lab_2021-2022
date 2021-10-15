/*
Calcolo dell'area di un poligono regolare

Dati il numero di lati n e la lunghezza di un lato l, il programma calcola l'area del poligono corrispondente
*/
package main

import (
	"fmt"
	"math"
)

func main() {

	var (
		n int     // numero di lati
		l float64 // lunghezza lato
	)

	fmt.Print("Inserisci il numero di lati del poligono: ")
	fmt.Scan(&n)
	fmt.Print("Inserisci la lunghezza dei lati del poligono: ")
	fmt.Scan(&l)

	var area float64
	// Pow() e Tan() sono due funzioni del package math che calcola rispettivamente l'elevamento a potenza e la tangente.
	// Pi è invece una costante del package math che rappresenta la costante matematica Pi greco
	area = (float64(n) * math.Pow(l, 2.0)) /
		(4 * math.Tan(math.Pi/float64(n)))
	// dato che il numero di lati del poligono è un numero intero, è necessario fare un cast esplicito da int a float64

	fmt.Println("Area calcolata:", area)
}
