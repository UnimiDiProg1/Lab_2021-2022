# Qual è l'output?

Qual è l'output del seguente programma?

```go
package main

import (
	"fmt"
)

func main() {

	var n int = 5
	var m int = 4

	var a [][]int

	a = make([][]int,n)

	for i := 0; i < n; i++ {
		a_i := make([]int,m)
		
		for j := i; j < m+i; j++ {
			a_i = append(a_i,j)
			fmt.Printf("%d ",j)
		}
		
		a = append(a, a_i)
		fmt.Println()
	}
}
```
