# 2 Qual è l'output?

Qual è l'output del seguente programma?

```go
package main

import (
	"fmt"
)

func main() {

	s := "Papà!"

	var s1 string
	for _, carattere := range s {
		s1 += string(carattere)
	}
	fmt.Println(s1)

	var s2 string
	for _, carattere := range s {
		s2 = string(carattere) + s2
	}
	fmt.Println(s2)

}
```