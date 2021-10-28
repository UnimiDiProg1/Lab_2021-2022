package main

import "fmt"

func main() {
	var dimensione int
	fmt.Print("Inserisci la dimensione: ")
	fmt.Scan(&dimensione)

	StampaScacchiera(dimensione)
}

func StampaRigaInizioAsterisco(lunghezza int) {
	var indiceColonna int
	for indiceColonna = 0; indiceColonna < lunghezza; indiceColonna++ {
		if indiceColonna%2 == 0 {
			fmt.Print("*")
		} else {
			fmt.Print("+")
		}
	}
}

func StampaRigaInizioPiù(lunghezza int) {
	var indiceColonna int
	for indiceColonna = 0; indiceColonna < lunghezza; indiceColonna++ {
		if indiceColonna%2 == 0 {
			fmt.Print("+")
		} else {
			fmt.Print("*")
		}
	}
}

func StampaScacchiera(dimensione int) {
	var indiceRiga int
	for indiceRiga = 0; indiceRiga < dimensione; indiceRiga++ {
		if indiceRiga%2 == 0 {
			StampaRigaInizioAsterisco(dimensione)
		} else {
			StampaRigaInizioPiù(dimensione)
		}
		fmt.Println()
	}
}
