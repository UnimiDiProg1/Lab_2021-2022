package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	fmt.Println("Inserisci un testo su più righe (termina con CTRL+D):")
	testo := LeggiTesto()
	testoSpaziato := Spazia(testo)

	fmt.Println("Risultato:")
	fmt.Println(testoSpaziato)
}

func LeggiTesto() (testo string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		testo += scanner.Text() + "\n"
	}
	return
}

func Spazia(testo string) (testoTrasformato string) {
	// similmente all'esercizio 3, la domanda che devo pormi è: quando devo inserire uno spazio?
	// la risposta è: ogni volta che una lettera era preceduta da un'altra lettera o, in generale,
	// quando un carattere diverso dallo spazio è preceduto da un carattere anch'esso diverso dallo spazio.
	// Per questo motivo uso una variabile booleana per memorizzare se il carattere precedente a quello attuale
	// era uno spazio. Se il carattere precedente non era uno spazio allora aggiungo uno spazio prima di concatenare
	// il carattere attuale.
	var spazioPrecedente bool = true
	for _, carattere := range testo {
		if unicode.IsSpace(carattere) {
			spazioPrecedente = true
		} else {
			if !spazioPrecedente {
				testoTrasformato += " "
			}
			spazioPrecedente = false
		}
		testoTrasformato += string(carattere)
	}
	return
}
