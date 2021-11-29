# Qual è l'output?

Qual è l'output di questo programma?

```go
package main

import "fmt"

func main() {

	var s string = "Ma il cielo è sempre più blu!"

	var s2 []rune = []rune(s)

	fmt.Println(len(s), len(s2))
}
```
