package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {

	fmt.Println("Inserisci un testo su più righe (termina con riga vuota):")
	testo := LeggiTesto()
	testoTrasformato := Garibaldi(testo)

	fmt.Println("Risultato trasformazione:")
	fmt.Println(testoTrasformato)

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
	return testo[:len(testo)-1]
}

func TrasformaCarattere(carattere rune) rune {
	if strings.ContainsRune("aeioàáèéìíùúòó", unicode.ToLower(carattere)) {
		if unicode.IsLower(carattere) {
			return 'u'
		} else {
			return 'U'
		}
	} else {
		return carattere
	}
}

func Garibaldi(testo string) (testoTrasformato string) {
	for _, l := range testo {
		testoTrasformato += string(TrasformaCarattere(l))
	}
	return
}
