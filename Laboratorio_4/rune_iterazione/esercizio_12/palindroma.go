package main

import "fmt"

func main() {
	var parola string
	fmt.Scan(&parola)

	if √ąPalindroma(parola) {
		fmt.Println("Palindroma")
	} else {
		fmt.Println("Non palindroma")
	}
}

func √ąPalindroma(parola string) bool {
	parolaCopia := ""
	for _, lettera := range parola {
		parolaCopia = string(lettera) + parolaCopia
	}
	return parolaCopia == parola
}
