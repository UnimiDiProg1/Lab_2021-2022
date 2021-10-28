# Qual è l'output?

Qual è l'output del seguente programma?

```go
package main

import "fmt"

var a int = 10

func test1() int {
	a += 5
	return a
}

func test2(a int) int {
	a += 6
	return a
}

func test3() int {
	return a + 7
}

func main() {
	var a, b, c int
	a, b, c = test1(), test2(a), test3()
	fmt.Println(a, b, c)
}
```