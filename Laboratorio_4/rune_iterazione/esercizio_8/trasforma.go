package main

import "fmt"

func inMaiuscolo(carattere rune) rune {
	if carattere >= 'a' && carattere <= 'z' {
		return 'A' + (carattere - 'a')
	}

	return carattere
}

func inMinuscolo(carattere rune) rune {
	if carattere >= 'A' && carattere <= 'Z' {
		return 'a' + (carattere - 'A')
	}

	return carattere
}

func main() {

	var testo string
	fmt.Scan(&testo)

	fmt.Print("Testo maiuscolo: ")
	for _, carattere := range testo {
		fmt.Printf("%c", inMaiuscolo(carattere))
	}
	fmt.Println()

	fmt.Print("Testo minuscolo: ")
	for _, carattere := range testo {

		fmt.Printf("%c", inMinuscolo(carattere))
	}
	fmt.Println()

}
