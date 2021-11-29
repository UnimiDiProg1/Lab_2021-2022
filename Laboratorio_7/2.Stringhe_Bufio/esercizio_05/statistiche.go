package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("Inserisci un testo su più righe (termina con Ctrl+D):")
	numeroLinee, numeroParole, lunghezza := StatisticheParole(LeggiTesto())

	fmt.Println("Statistiche:")
	fmt.Println("Numero linee:", numeroLinee)
	fmt.Println("Numero parole:", numeroParole)
	fmt.Println("Lunghezza media:", float64(lunghezza)/float64(numeroParole))
}

func LeggiTesto() (testo string) {
	// leggo del testo dallo standard input (tastiera) e lo concateno nella variabile testo
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		testo += scanner.Text() + "\n"
	}
	return
}

func ContaLettere(parola string)(numeroLettere int){

	for _,v := range parola{

		if unicode.IsLetter(v){

			numeroLettere++

		}

	}

	return

}

func StatisticheParole(testo string) (numeroLinee int, numeroParole int, lunghezzaTotale int) {

	//la funzione Split permette di ottenere una slice di stringe
	//ogni stringa nello slice è una linea
	for _, linea := range strings.Split(testo, "\n") {

		//ignoro le linee vuote
		if len(linea) > 0 {
			numeroLinee++
		}

		//qui Split viene usata per separare le sequenze di caratteri intervallati da uno spazio
		//l'output è una slice di stringhe
		for _, parola := range strings.Split(linea, " ") {
			//ignoro le stringhe vuote
			if len(parola) > 0 {
				numeroParole++
				lunghezzaTotale += ContaLettere(parola)
			}
		}

	}

	return
}
