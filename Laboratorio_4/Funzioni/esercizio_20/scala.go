package main

import (
	"fmt"
)

const PROFONDITÀ_GRADINO = 3

func main() {

	var numeroGradini int
	fmt.Print("Inserisci il numero dei gradini: ")
	fmt.Scan(&numeroGradini)

	StampaScala(numeroGradini)
}

func StampaRientro(gradino int) {
	var dimensioneRientro int = gradino * (PROFONDITÀ_GRADINO - 1)
	var contatore int
	for contatore = 0; contatore < dimensioneRientro; contatore++ {
		fmt.Print(" ")
	}
}

func StampaGradino(gradino int) {
	if gradino >= 0 {
		StampaRientro(gradino)

		var contatore int
		for contatore = 0; contatore < PROFONDITÀ_GRADINO; contatore++ {
			fmt.Print("*")
		}
		fmt.Println()

		StampaRientro(gradino)

		fmt.Println("*")
	}
}

func StampaScala(gradini int) {
	if gradini > 0 {
		var gradino int
		for gradino = gradini - 1; gradino >= 0; gradino-- {
			StampaGradino(gradino)
		}
	} else {
		fmt.Println("Dimensione non sufficiente")
	}
}
