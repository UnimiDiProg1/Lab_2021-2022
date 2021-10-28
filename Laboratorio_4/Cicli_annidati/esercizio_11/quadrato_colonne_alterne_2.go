package main

import (
	"fmt"
)

func main() {

	var dimensione int
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&dimensione)

	// stampa n righe di coppie di caratteri * e + alternati sulla stessa riga
	var indiceRiga, indiceColonna int
	for indiceRiga = 0; indiceRiga < dimensione; indiceRiga++ {
		for indiceColonna = 0; indiceColonna < dimensione; indiceColonna++ {
			// Un modo per risolvere questo problema è quello di considerare 4 casi distinti:
			// caso 1 e 2 -> stampa di un *
			// caso 3 e 4 -> stampa di un +
			// I casi sono 4 perché il pattern che si ripete è di 4 caratteri (**++)

			// utilizzo il resto della divisione tra l'indice di colonna e 4 perché ho 4 casi
			if indiceColonna%4 <= 1 {
				// i primi due casi corrispondono ad un resto uguale a 0 o a 1
				fmt.Print("* ")
			} else {
				// i rimanenti due casi corrispondono ad un resto uguale a 2 o a 3
				fmt.Print("+ ")
			}
		}
		fmt.Println()
	}

}
