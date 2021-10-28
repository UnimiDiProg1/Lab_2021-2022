# Qual è l'output?

Quanti asterischi stampa il seguente programma?

```go
package main

import "fmt"

func main() {
	var n int = 3

	var i, j, z int
	for i = 0; i < n; i++ {
		for j = 0; j < n; j++ {
			for z = 0; z < n; z++ {
				fmt.Print("*")
			}
		}
	}
	fmt.Println()
}
```
