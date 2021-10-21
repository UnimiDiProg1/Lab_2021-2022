package main

import (
	"fmt"
	"math"
)

func main() {
	// Tute le variabili intere sono inizializzate a 0 di default
	var (
		n, // il numero letto
		min, max, // variabili per memorizzare massimo e minimo
		somma, // memorizza la somma
		positivi, negativi, nulli int // memorizzano il numero di numeri positivi, negativi e nulli inseriti
	)

	fmt.Scan(&n)
	fmt.Print("Inserisci ", n, " numeri\n")

	// Inizializzazione di minimo e massimo
	min, max = math.MaxInt64, math.MinInt64
	// Per inizializzare il valore di min e max si possono usare un valore molto alto e un valore molto basso.
	// In questo modo i numeri letti successivamente dovrebbero permettere la loro modifica.
	// Dato che però non è possibile sapere a priori che numeri saranno inseriti, min e max dovrebbero essere inizializzati così:
	// min, max = math.MaxInt64, math.MinInt64
	// ovvero con il valore intero massimo e il valore intero minimo rappresentabili

	var i int;
	for i = 0; i < n; i++ {
		// Ad ogni iterazione leggo un numero e aggiorno i valori

		var numeroLetto int
		fmt.Scan(&numeroLetto)

		// Incremento somma
		somma += numeroLetto

		// Min e max
		if numeroLetto > max {
			max = numeroLetto
		}
		if numeroLetto < min {
			min = numeroLetto
		}

		// Test segno
		if numeroLetto > 0 {
			positivi++
		} else if numeroLetto < 0 {
			negativi++
		} else {
			nulli++
		}
	}

	fmt.Print("somma = ", somma, "\n")
	fmt.Print("valore minimo = ", min, "\n")
	fmt.Print("valore massimo = ", max, "\n")
	fmt.Print("interi > 0 = ", positivi, "\n")
	fmt.Print("interi < 0 = ", negativi, "\n")
	fmt.Print("interi = 0 = ", nulli, "\n")

}
