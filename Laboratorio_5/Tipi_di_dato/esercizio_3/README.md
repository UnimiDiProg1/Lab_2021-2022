# Trova l'errore

Questo programma dovrebbe stampare il valore delle variabili `a`, `b`, `c` e `d` ma contiene degli errori. Corregere gli errori e verificare che l'esecuzione produca l'output desiderato.


```go
package main

import "fmt"

func main() {
    var a int = 4
    var b float64 = 12.5
    var c int
    c := a + b
	
    var d float64
    d = a/c
	
    fmt.Println(a, b, c, d)
}
```