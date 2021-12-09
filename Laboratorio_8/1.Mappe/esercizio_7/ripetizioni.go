package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {

	ParoleRipetute(LeggiTesto())

}

func LeggiTesto() (testo string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		testo += scanner.Text() + "\n"
	}
	return
}

func SeparaParole(testo string) (parole []string) {
	parola := ""
	for _, carattere := range testo {
		if unicode.IsLetter(carattere) {
			parola += string(carattere)
		} else if parola != "" {
			parole = append(parole, parola)
			parola = ""
		}
	}
	return
}

func ContaRipetizioni(testo string) (ripetizioni map[string]int) {
	ripetizioni = map[string]int{}
	for _, parola := range SeparaParole(testo) {
		ripetizioni[parola]++
	}
	return
}

func ParoleRipetute(testo string) {
	ripetizioni := ContaRipetizioni(testo)
	fmt.Printf("Parole distinte: %d\n", len(ripetizioni))
	for parola, conteggio := range ripetizioni {
		fmt.Printf("%s: %d\n", parola, conteggio)
	}
	fmt.Println()
}
