# Qual è l'output?

Qual è l'output di questo programma?

```go
package main

import "fmt"

func main() {
	var a [6]int

	for i := range a {
		a[i] = i
	}

	for _, v := range a {
		v *= 2
	}

	fmt.Println(a)
}
```
