package main

import "fmt"
import "math"
import "math/rand"

const EPSILON = 1.0e-9

func main() {

	var (
		s int64
		m, q float64
	)

	fmt.Print("Inserisci s: ")
	fmt.Scan(&s)

	fmt.Print("Inserisci m e q: ")
	fmt.Scan(&m, &q)

	fmt.Println()

	rand.Seed(s)
	for i:=0;i<10;i++{
		x := rand.Float64() * 5.0
		y := rand.Float64() * 5.0
		fmt.Print("Punto (",x,",",y,") - ")
		if ÈXMaggioreDiY(y,m*x+q) {
			fmt.Println("Il punto sta sopra la retta")
		} else if ÈXMinoreDiY(y,m*x+q) {
			fmt.Println("Il punto sta sotto la retta")
		} else {
			fmt.Println("Il punto appartiene alla retta")
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
