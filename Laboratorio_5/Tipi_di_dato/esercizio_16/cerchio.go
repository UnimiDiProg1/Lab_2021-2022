package main

import "fmt"
import "math"
import "math/rand"

const EPSILON = 1.0e-6

func main() {

	var (
		s int64
		xC, yC float64
		raggio float64
	)

	fmt.Print("Inserisci s: ")
	fmt.Scan(&s)

	fmt.Print("Inserisci i valori dell'ascissa e dell'ordinata del punto C (il centro del cerchio): ")
	fmt.Scan(&xC, &yC)

	fmt.Print("Inserisci il valore del raggio: ")
	fmt.Scan(&raggio)

	fmt.Println()

	delta := 2.0 * raggio
	xOffset := xC - 1.0 * raggio
	yOffset := yC - 1.0 * raggio

	rand.Seed(s)

	xMin, yMin := xC, yC + raggio // un punto sulla circonferenza
	distanzaMin := raggio

	xMax, yMax := xC, yC + raggio // un punto sulla circonferenza
	distanzaMax := raggio

	for i:=0;i<1000000;i++{

		x := xOffset + (rand.Float64() * delta)
		y := yOffset + (rand.Float64() * delta)

		distanza := Distanza(x, y, xC, yC)

		//fmt.Print("Punto (", x, ",", y, ").\n")

		if ÈXUgualeAY(distanza, raggio) {
			fmt.Print("Il punto (", x, ",", y, ") giace sulla circonferenza del cerchio.\n")
		} else if ÈXMaggioreDiY(distanza, raggio) {
			if ÈXMaggioreDiY(distanza, distanzaMax) {
				distanzaMax = distanza
				xMax = x
				yMax = y
			}
		} else {
			if ÈXMinoreDiY(distanza, distanzaMin) {
				distanzaMin = distanza
				xMin = x
				yMin = y
			}
		}
	}

	if ÈXUgualeAY(distanzaMax, raggio) {
		fmt.Println("\nNon esistono punti esterni al cerchio tra tutti quelli generati.")
	} else {
		fmt.Print("\nIl punto (", xMax, ",", yMax, ") è quello all'esterno del cerchio che ha distanza massima dal centro C.\n")
		fmt.Print("Distanza: ",distanzaMax,"\n")
	}

	if ÈXUgualeAY(distanzaMin, raggio) {
		fmt.Println("\nNon esistono punti interni al cerchio tra tutti quelli generati.")
	} else {
		fmt.Print("\nIl punto (", xMin, ",", yMin, ") è quello all'interno del cerchio che ha distanza minima dal centro C.\n")
		fmt.Print("Distanza: ",distanzaMin,"\n")
	}

}

func Distanza(x1, y1 float64, x2, y2 float64) float64 {

	return math.Sqrt((x1-x2) * (x1-x2) + (y1-y2) * (y1-y2))

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
	return math.Abs(x - y) <= EPSILON
}

/* La funzione `func ÈXMaggioreDiY(x, y float64) bool` riceve in input due
valore reali nei parametri `x` e `y` e restituisce `true` se `x` è minore
di `y` di almeno una costante EPSILON */
func ÈXMinoreDiY(x, y float64) bool {
	return (x - y) < -EPSILON
}
