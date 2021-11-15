package main

import "fmt"
import "math"
import "math/rand"

const EPSILON = 1.0e-9

func main() {

	var (
		s int64
		m1, q1 float64
		m2, q2 float64
		m3, q3 float64
	)

	fmt.Print("Inserisci s: ")
	fmt.Scan(&s)

	fmt.Print("Inserisci m1 e q1: ")
	fmt.Scan(&m1, &q1)

	fmt.Print("Inserisci m2 e q2: ")
	fmt.Scan(&m2, &q2)

	fmt.Print("Inserisci m3 e q3: ")
	fmt.Scan(&m3, &q3)

	fmt.Println()

	rand.Seed(s)
	for i:=0;i<10;i++{
		x := rand.Float64() * 10.0
		y := rand.Float64() * 10.0
		fmt.Print("Punto (",x,",",y,") - ")
		if ÈXMinoreDiY(y,m1*x+q1) || ÈXMaggioreDiY(y,m2*x+q2) || ÈXMaggioreDiY(y,m3*x+q3)  {
			fmt.Println("Il punto sta all'esterno del triangolo.")
		} else {
			fmt.Println("Il punto sta all'interno del triangolo.")
		}
	}
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
