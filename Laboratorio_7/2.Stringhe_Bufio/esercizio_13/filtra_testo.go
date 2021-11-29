package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	testoInRighe := LeggiTesto()
	fmt.Println("Testo filtrato:")
	for _, riga := range FiltraTesto(testoInRighe) {
		fmt.Println(riga)
	}
}

func LeggiTesto() (testo []string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		testo = append(testo, scanner.Text())
	}
	return
}

func ContieneCifre(testo string) bool {
	for _, carattere := range testo {
		if unicode.IsDigit(carattere) {
			return true
		}
	}
	return false
}

func FiltraTesto(testo []string) (testoFiltrato []string) {
	for _, riga := range testo {
		if ContieneCifre(riga) {
			testoFiltrato = append(testoFiltrato, riga)
		}
	}
	return
}
