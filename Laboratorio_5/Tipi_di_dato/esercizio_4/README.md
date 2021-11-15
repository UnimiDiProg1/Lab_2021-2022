# Valore di default delle variabili

Che cosa succede se si prova a utilizzare una variabile senza averla inizializzata? Si esegua il seguente programma per stampare a video i valori di default (valori iniziali/zero-values) per i tipi `int`, `float64`, `string` e `bool`.

```go
package main

import "fmt"

func main() {

	var numeroIntero int
	var numeroReale float64
	var valoreLogico bool

	fmt.Print("Valore di default per il tipo int: _", numeroIntero, "_\n")
	fmt.Print("Valore di default per il tipo float64: _", numeroReale, "_\n")
	fmt.Print("Valore di default per il tipo bool: _", valoreLogico, "_\n")

}
```