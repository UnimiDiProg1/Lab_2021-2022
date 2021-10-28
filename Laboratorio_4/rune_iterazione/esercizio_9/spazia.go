package main

import "fmt"

func main() {

	var testo string
	fmt.Print("Inserisci una stringa di testo: ")
	fmt.Scan(&testo)

	for _, c := range testo {
		fmt.Printf("%c ", c)
	}
	fmt.Println()

}
