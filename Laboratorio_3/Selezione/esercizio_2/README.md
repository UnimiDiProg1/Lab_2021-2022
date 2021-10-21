# Qual è l'output?

Supponendo che l'utente inserisca da **standard input**
```
5 8 6
```
qual è l'output del seguente programma? Che cosa calcola? 

```go
package main

import "fmt"

func main() {

	var a, b, c int
	var m int

	fmt.Scan(&a, &b, &c)

	if a < b {
		if a < c {
			m = a
		} else {
			m = c
		}
	} else {
		if b < c {
			m = b
		} else {
			m = c
		}
	}

	fmt.Println(m)

}
```
