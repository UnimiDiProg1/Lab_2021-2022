# Qual è l'output?

Qual è l'output di questo programma?

```go
package main

import "fmt"

func main() {

	var n int = 5

	var s []int
	s = make([]int, n)

	for i := 0; i < n; i++ {
		s[i] = i
	}

	fmt.Println(s)
}
```
