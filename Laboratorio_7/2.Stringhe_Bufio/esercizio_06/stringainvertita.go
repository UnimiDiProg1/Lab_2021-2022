package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Inserisci un testo su più righe (termina con riga vuota):")
	testo := LeggiTesto()
	testoInvertito := InvertiStringa(testo)

	fmt.Println("Testo invertito:")
	fmt.Println(testoInvertito)

}

func LeggiTesto() (testo string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		riga := scanner.Text()
		if riga != "" {
			testo += riga + "\n"
		} else {
			break
		}
	}

	// In questo modo restituisco il testo letto in input senza l'ultimo carattere (il carattere '\n')
	return testo[:len(testo)-1]
}

func InvertiStringa(stringa string) (stringaOutput string) {
	for _, c := range stringa {
		stringaOutput = string(c) + stringaOutput
	}
	return
}
