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

		for contatore = 0; contatore < PROFONDITÀ_GRADINO-1; contatore++ {
			fmt.Print(" ")
		}
		fmt.Println("*")
	}
}

func StampaScala(gradini int) {
	if gradini > 0 {
		var gradino int
		for gradino = 0; gradino < gradini; gradino++ {
			StampaGradino(gradino)
		}
	} else {
		fmt.Println("Dimensione non sufficiente")
	}
}
