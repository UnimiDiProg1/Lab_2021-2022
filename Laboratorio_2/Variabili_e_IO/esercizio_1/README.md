# Scan

Il programma riportato di seguito legge da **standard input** 4 numeri interi.

* Cosa succede se inserite l'input su righe diverse?
* Cosa succede se inserite meno numeri di quelli richiesti?
* Cosa succede se inserite più numeri di quelli richiesti?
* Cosa succede se inserite un valore diverso da un numero intero (numero reale, lettera, parola, ...)?

```go
package main

import "fmt"

func main() {
	
	var n1, n2, n3, n4 int

	fmt.Scan(&n1, &n2, &n3, &n4)

	fmt.Println(n1, n2, n3, n4)
}
```
