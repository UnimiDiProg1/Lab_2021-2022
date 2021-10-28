package main

import "fmt"

func main() {

	var numero int
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&numero)
	Tabellina(numero)

}

func Tabellina(numero int) {
	if numero >= 1 && numero <= 9 {
		fmt.Print("Tabellina del ", numero, ":")
		var contatore int
		for contatore = 0; contatore <= 10; contatore++ {
			fmt.Print(" ", contatore*numero)
		}
		fmt.Println()
	}
}
