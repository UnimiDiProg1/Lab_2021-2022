# Qual è l'output?

Supponendo che l'utente inserisca da **standard input** `36`, qual è l’output di questo programma? 
E se inserisse `32`?
Che cosa fa questo programma? Quali sono le differenze con il programma dell'esercizio precedente?

```go
package main

import "fmt"

func main() {

	var n int

	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&n)

	var i int;
	for i = 1; i <= n; i *= 2 {
		fmt.Print(i, " ")
	}

	fmt.Println()

}
```
