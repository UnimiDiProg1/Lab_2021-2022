package main

import (
	"fmt"
	"unicode"
)

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

		if unicode.IsLetter(l) {
			if unicode.IsUpper(l) {
				if èVocale(l) {
					vmaiuscole++
				} else {
					cmaiuscole++
				}
			} else {
				if l == 'a' || l == 'e' || l == 'i' || l == 'o' || l == 'u' {
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
