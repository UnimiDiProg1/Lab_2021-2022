package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	// leggo una singola riga di testo
	if scanner.Scan() {
		testo := scanner.Text()
		numero, err := NumeroNascosto(testo)
		// se la ricerca del numero nascosto è andata a buon fine stampo in output il numero nascosto
		if err == nil {
			fmt.Print("Doppio del numero nascosto: ", 2*numero, " (", numero, " * 2)\n")
		}
	}
}

func NumeroNascosto(testo string) (int, error) {
	// la funzione si divide in due fasi:

	// fase 1: ricerca del numero nascosto, ovvero la ricerca di tutti i caratteri che rappresentano delle
	// cifre all'interno del testo fornito in input
	numeroNascosto := ""
	for _, carattere := range testo {
		if unicode.IsDigit(carattere) {
			numeroNascosto += string(carattere)
		}
	}

	// fase 2: coversione del testo che contiene il numero nascosto da stringa ad intero.
	// La funzione strconv.Atoi restituisce due valori: un numero intero e un eventuale errore occorso durante la
	// conversione
	return strconv.Atoi(numeroNascosto)
}
