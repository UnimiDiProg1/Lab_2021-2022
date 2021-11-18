# Qual è l'output?

Qual è l'output di questo programma?

```go
package main

import "fmt"

func main() {
	var a = [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(a)
	modifica(a)
	fmt.Println(a)
}

func modifica(a [6]int) {
	for i := range a {
		a[i] *= 2
	}
}
```
