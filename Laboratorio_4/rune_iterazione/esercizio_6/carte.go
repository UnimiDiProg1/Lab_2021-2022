package main

import "fmt"

func main() {

	// Codice unicode che corrisponde all'asso di cuori
	const ACE = '\U0001F0B1'

	// Codice unicode che corrisponde al 10 di cuori, ovvero il codice dell'asso + 9
	const TEN = '\U0001F0B1' + 9

	// Per tutte i codice compresi tra l'asso e il 10, stampp il carattere corrispondente e il codice numerico
	// I codici tra ACE e TEN sono i codici delle carte di cuori
	for c := ACE; c <= TEN; c++ {
		fmt.Println("Simbolo:", string(c), "- Codice numerico in base 10:", c)
	}

}
