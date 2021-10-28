package main

import "fmt"

func main() {
	var n int
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&n)

	if ÈPerfetto(n) {
		fmt.Println(n, "è perfetto")
	} else {
		fmt.Println(n, "non è perfetto")
	}

}

func SommaDivisoriPropri(n int) (somma int) {
	var divisore int
	for divisore = 1; divisore < n; divisore++ {
		if n%divisore == 0 {
			somma += divisore
		}
	}
	return
}

func ÈPerfetto(n int) bool {
	return n > 0 && SommaDivisoriPropri(n) == n
}
