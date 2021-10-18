# Trova gli errori

Questo programma è errato e non produce l’output descritto nel commento del package. Correggilo (3 errori: 1 di sintassi,
2 di logica del programma).

```go
/*
Il programma stampa una sequenza di numeri

Dato un numero n inserito da tastiera, il programma stampa tutti i numeri 
compresi nell'intervallo che va da 1 a n (estremi inclusi).
La sequenza è prodotta su un'unica riga di testo in cui ciascun numero è separato
dal precedente da uno spazio.
*/
package main

import (
	"fmt"
)

func main() {

	var n int

	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&n)

	var i int;
	for i = 1; i < n { 
		fmt.Print(i)
		i++
	}

	fmt.Println()

}
```
