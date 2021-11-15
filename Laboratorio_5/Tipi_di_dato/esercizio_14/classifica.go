package main

import (
	"fmt"
	"math"
)

const EPSILON = 1.0e-9

func main() {

	var (
		xA, yA float64
		xB, yB float64
		xC, yC float64
	)

	fmt.Print("Inserisci i valori dell'ascissa e dell'ordinata del punto A: ")
	fmt.Scan(&xA, &yA)

	fmt.Print("Inserisci i valori dell'ascissa e dell'ordinata del punto B: ")
	fmt.Scan(&xB, &yB)

	fmt.Print("Inserisci i valori dell'ascissa e dell'ordinata del punto C: ")
	fmt.Scan(&xC, &yC)

	fmt.Println()

	lunghezzaAB := Distanza(xA, yA, xB, yB)
	lunghezzaAC := Distanza(xA, yA, xC, yC)
	lunghezzaBC := Distanza(xB, yB, xC, yC)

	switch {
	case ÈXUgualeAY(lunghezzaAB, lunghezzaAC) && ÈXUgualeAY(lunghezzaAB, lunghezzaBC):
		DescriviTriangoloEquilatero(lunghezzaAC)
	case ÈXUgualeAY(lunghezzaAB, lunghezzaAC) || ÈXUgualeAY(lunghezzaAB, lunghezzaBC) || ÈXUgualeAY(lunghezzaBC, lunghezzaAC):
		DescriviTriangoloIsoscele(lunghezzaAB, lunghezzaBC, lunghezzaAC)
	default:
		DescriviTriangoloScaleno(lunghezzaAB, lunghezzaBC, lunghezzaAC)
	}
}

/*
La funzione `Distanza(x1, y1 float64, x2, y2 float64) float64` riceve in input:

    - due valori `float64` nei parametri `x1` e `y1` che rappresentano rispettivamente l'ascissa e l'ordinata di un punto `P1` sul piano cartesiano;

    - due valori `float64` nei parametri `x2` e `y2` che rappresentano rispettivamente l'ascissa e l'ordinata di un punto `P2` sul piano cartesiano;

    e restituisce un valore `float64` pari alla distanza euclidea tra i punti `P1` e `P2`.

*/
func Distanza(x1, y1, x2, y2 float64) float64 {

	return math.Sqrt((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2))

}

/* La funzione `func ÈXMaggioreDiY(x, y float64) bool` riceve in input due
valore reali nei parametri `x` e `y` e restituisce `true` se `x` è maggiore
di `y` di almeno una costante EPSILON */
func ÈXMaggioreDiY(x, y float64) bool {
	return (x - y) > EPSILON
}

/* La funzione `func ÈXUgualeAY(x, y float64) bool` riceve in input due
valore reali nei parametri `x` e `y` e restituisce `true` se `x` è uguale
a `y` a meno di una costante EPSILON */
func ÈXUgualeAY(x, y float64) bool {
	return math.Abs(x-y) <= EPSILON
}

/* La funzione `func ÈXMaggioreDiY(x, y float64) bool` riceve in input due
valore reali nei parametri `x` e `y` e restituisce `true` se `x` è minore
di `y` di almeno una costante EPSILON */
func ÈXMinoreDiY(x, y float64) bool {
	return (x - y) < -EPSILON
}

func DescriviTriangoloEquilatero(lunghezzaLato float64) {
	fmt.Println("Il triangolo ABC è equilatero.")
	fmt.Println("Lunghezza del lato:", lunghezzaLato)
}

func DescriviTriangoloIsoscele(lunghezzaAB, lunghezzaBC, lunghezzaAC float64) {
	fmt.Println("Il triangolo ABC è isocele.")
	if ÈXUgualeAY(lunghezzaAB, lunghezzaAC) {
		fmt.Println("I lati di lunghezza uguale sono AB e AC.")
		fmt.Println("Lunghezza dei lati AB e AC:", lunghezzaAB)
		fmt.Println("Lunghezza del lato BC:", lunghezzaBC)
	} else if ÈXUgualeAY(lunghezzaAB, lunghezzaBC) {
		fmt.Println("I lati di lunghezza uguale sono AB e BC.")
		fmt.Println("Lunghezza dei lati AB e BC:", lunghezzaAB)
		fmt.Println("Lunghezza del lato AC:", lunghezzaAC)
	} else {
		fmt.Println("I lati di lunghezza uguale sono AC e BC.")
		fmt.Println("Lunghezza dei lati AC e BC:", lunghezzaAC)
		fmt.Println("Lunghezza del lato AB:", lunghezzaAB)
	}
}

func DescriviTriangoloScaleno(lunghezzaAB, lunghezzaBC, lunghezzaAC float64) {
	fmt.Println("Il triangolo ABC è scaleno.")
	fmt.Println("Lunghezza del lato AB:", lunghezzaAB)
	fmt.Println("Lunghezza del lato BC:", lunghezzaBC)
	fmt.Println("Lunghezza del lato AC:", lunghezzaAC)
}
