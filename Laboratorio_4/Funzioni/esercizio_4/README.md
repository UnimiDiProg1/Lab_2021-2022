# Trova l'errore

Questo programma dovrebbe stampare `20 40` ma non genera l'output desiderato.
Corregere e verificare che l'esecuzione sia conforme al comportamento atteso.

```go
package main

import "fmt"

func test(x, y int) (int, int) {
	return 2 * x, 2 * y
}

func main() {
	var x, y int = 10, 20
	test(x, y)
	fmt.Println(x, y)
}
```