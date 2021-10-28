package main

import (
	"fmt"
)

func èLetteraValida(l rune) bool {
	if (l >= 'A' && l <= 'Z') || (l >= 'a' && l <= 'z') {
		return true
	}

	return false
}

func èMaiuscola(l rune) bool {
	if l >= 'A' && l <= 'Z' {
		return true
	}

	return false
}

func èVocale(l rune) bool {
	switch l {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	default:
		return false
	}
}

func main() {

	maiuscole, minuscole := 0, 0
	consonanti, vocali := 0, 0

	var stringaInput string

	fmt.Scan(&stringaInput)

	for _, l := range stringaInput {

		if èLetteraValida(l) {
			if èMaiuscola(l) {
				maiuscole++
				l = l - 'A' + 'a'
			} else {
				minuscole++
			}

			if èVocale(l) {
				vocali++
			} else {
				consonanti++
			}
		}

	}
	fmt.Println("Maiuscole:", maiuscole)
	fmt.Println("Minuscole:", minuscole)
	fmt.Println("Vocali:", vocali)
	fmt.Println("Consonanti:", consonanti)
}
