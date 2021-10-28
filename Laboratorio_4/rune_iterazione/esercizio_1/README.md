# Qual è l'output?

Qual è l'output del seguente programma?

```go
package main

import "fmt"

func main() {
	var x int
	var y int = 'a'
	x = 'A'
	fmt.Print(x, " ")
	for x := 1; x < 10; x++ {
		fmt.Print(x+y, " ")
	}
	fmt.Println()
}
```