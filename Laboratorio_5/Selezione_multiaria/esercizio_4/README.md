# Qual è l'output?

Supponendo che l'utente inserisca da **standard input** il valore `3`, cosa dovrebbe produrre in output il seguente programma? E se invece inserisse `100`?

```go
package main

import "fmt"

func main() {

	var n int

	fmt.Scan(&n)

	accumulatore := 1

	for i:=1; i<=n; i++ {

		switch {
		case i<5:
			accumulatore *= i
		case i<10:
			accumulatore += i
		}

	}

	fmt.Println("Accumulatore =", accumulatore)
}
```
 