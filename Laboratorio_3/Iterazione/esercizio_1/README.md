# Qual è l'output?

Supponendo che l'utente inserisca da **standard input** `10`, qual è l’output di questo programma? 
E se inserisse `11`?
Che cosa fa questo programma?

```go
package main

import "fmt"

func main() {
	var n int
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&n)

	var i int;
	for i = 0; i <= n; i += 2 {
		fmt.Print(i, " ")
	}
	fmt.Println()
}
```
