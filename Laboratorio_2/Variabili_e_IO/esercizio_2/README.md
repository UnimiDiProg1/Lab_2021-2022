# Print o Println

Di seguito sono riportati due programmi: in entrambi sono dichiarate 2 variabili `a` e `b` di cui vengono stampati i valori.
Notate differenze nell'output prodotto? Cosa cambia tra `Print` e `Println`?

```go 
package main

import "fmt"

func main() {
	var a int = 5
	var b float64 = 3.14

	fmt.Print("Valore di a:", a, "capito? Te lo dico due volte:", a, a, "...\n")
	fmt.Print("Valore di b:", b, b, "\n")
}
```

```go
package main

import "fmt"

func main() {
	var a int = 5
	var b float64 = 3.14

	fmt.Println("Valore di a:", a, "capito? Te lo dico due volte:", a, a, "...")
	fmt.Println("Valore di b:", b, b)
}
```
