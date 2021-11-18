# Qual è l'output?

Qual è l'output di questo programma?

```go
package main

import "fmt"

func main() {
	var a = [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(a[:])

	b := a[:]
	b = modifica(b)

	fmt.Println(b)
}

func modifica(slIn []int) (slOut []int) {
    slOut = slIn
	for i := range slOut {
		slOut[i] *= 2
	}
	return
}
```
