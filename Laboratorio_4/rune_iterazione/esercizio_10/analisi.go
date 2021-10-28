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
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}
}

func main() {

	vmaiuscole, vminuscole := 0, 0
	cmaiuscole, cminuscole := 0, 0

	var stringaInput string

	fmt.Scan(&stringaInput)

	for _, l := range stringaInput {

		if èLetteraValida(l) {
			if èMaiuscola(l) {
				if èVocale(l) {
					vmaiuscole++
				} else {
					cmaiuscole++
				}
			} else {
				if èVocale(l) {
					vminuscole++
				} else {
					cminuscole++
				}
			}

		}

	}
	fmt.Println("Vocali maiuscole:", vmaiuscole)
	fmt.Println("Consonanti maiuscole:", cmaiuscole)
	fmt.Println("Vocali minuscole:", vminuscole)
	fmt.Println("Consonanti minuscole:", cminuscole)
}
