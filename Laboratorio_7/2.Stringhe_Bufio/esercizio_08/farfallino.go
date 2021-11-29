package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("Inserisci un testo su più righe (termina con CTRL+D):")
	testo := LeggiTesto()
	testoTrasformato := TraduciTestoInFarfallino(testo)

	fmt.Println("Risultato:")
	fmt.Println(testoTrasformato)

}

func LeggiTesto() (testo string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		testo += scanner.Text() + "\n"
	}
	return
}

func TraduciCarattereInFarfallino(carattere rune) (carattereTrasformato string) {
	// converto il carattere in una stringa a cui poi concatenerò (se necessario) la lettera 'f' e il carattere stesso
	carattereTrasformato = string(carattere)
	// invece di usare degli if o uno switch ho utilizzato la funzione strings.ContainsRune:
	// questa funzione restituisce true se all'interno della stringa passata come primo argomento è presente il
	// carattere passato come secondo argomento. È un metodo più veloce per controllare se un carattere rientra tra
	// le vocali.
	if strings.ContainsRune("aeiouèéòóùúàáíì", carattere) {
		// questo gestisce il caso in cui il carattere sia una vocale minuscola
		carattereTrasformato += "f" + carattereTrasformato
	}
	if strings.ContainsRune("AEIOUÀÁÈÉÌÍÓÒÚÙ", carattere) {
		// questo gestisce il caso in cui il carattere sia una vocale maiuscola
		carattereTrasformato += "F" + carattereTrasformato
	}
	return
}

func TraduciTestoInFarfallino(testo string) (testoTrasformato string) {
	for _, carattere := range testo {
		testoTrasformato += TraduciCarattereInFarfallino(carattere)
	}
	return
}
