package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {

	fmt.Print("Inserisci un numero: ")
	chiave := LeggiNumero()
	fmt.Println("Inserisci un testo su più righe (termina con CTRL D):")
	testo := LeggiTesto()

	testoCifrato := CifraTesto(testo, chiave)

	fmt.Println("Testo cifrato:")
	fmt.Println(testoCifrato)
}

func LeggiTesto() (testo string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		testo += scanner.Text() + "\n"
	}
	return
}

func LeggiNumero() (n int) {
	fmt.Scan(&n)
	return
}

func CifraCarattere(carattere rune, chiave int) rune {

	// il numero di lettere dell'alfabeto inglese.
	// se non me lo ricordo a memoria, posso calcolarlo come differenza tra l'ultima e la prima lettera
	const rangeLettere int32 = 'Z' - 'A' + 1

	var chiavePositiva rune
	// se la chiave è maggiore o uguale al numero di lettere dell'alfabeto è come se stessi chiedendo di fare più "giri"
	// dell'alfabeto: supponendo di avere una lettera 'A' e una chiave=25, la lettera risultante dalla traslazione
	// sarebbe 'Z'. Supponendo di avere una lettera 'A' e una chiave=26, la lettera risultante sarebbe la 'A' stessa,
	// ovvero sarei arrivato alla fine dell'alfabeto e ripartito da capo. Una chiave=26 equivale ad una chiave=0,
	// una chiave=27 equivale ad una chiave=1, e così via.
	chiavePositiva = int32(chiave) % rangeLettere

	// se la chiave dovesse essere negativa la rendo positiva: supponendo di avere una lettera 'A' e una chiave=-1
	// dalla codifica otterrei la lettera 'Z', che equivale al risultato che otterrei con chiave-25 (26 lettere - 1).
	if chiavePositiva < 0 {
		chiavePositiva = rangeLettere + chiavePositiva
	}

	// Se la lettera è una lettera dell'alfabeto inglese
	if (carattere >= 'A' && carattere <= 'Z') || (carattere >= 'a' && carattere <= 'z') {
		// imposto il mio carattere di partenza (la 'A' o la 'a' dipendentemente dal fatto che il carattere di
		// partenza sia maiuscolo o minuscolo)
		base := 'A'
		if unicode.IsLower(carattere) {
			base = 'a'
		}

		return (((carattere - base) + chiavePositiva) % (rangeLettere)) + (base)
	} else {
		return carattere
	}
}

func CifraTesto(testo string, chiave int) (testoCifrato string) {
	for _, carattere := range testo {
		testoCifrato += string(CifraCarattere(carattere, chiave))
	}
	return
}
