package main

import "fmt"

func main() {
	var parola string
	fmt.Scan(&parola)

	if ÈPalindroma(parola) {
		fmt.Println("Palindroma")
	} else {
		fmt.Println("Non palindroma")
	}
}

func ÈPalindroma(parola string) bool {
	parolaCopia := ""
	for _, lettera := range parola {
		parolaCopia = string(lettera) + parolaCopia
	}
	return parolaCopia == parola
}
