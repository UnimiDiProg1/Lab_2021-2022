# Qual è l'output?

Qual è l'output del seguente programma?

```go
package main

import "fmt"

func funzione1(x int) int {
	return x * 10
}

func funzione2(x int) (y int) {
	y = x * 10
	return
}

func main() {
	fmt.Println(funzione1(5), funzione2(3))
}
```