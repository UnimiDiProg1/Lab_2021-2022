# Qual è l'errore?

Si consideri il seguente frammento di codice. 

```go
package main

import "fmt"

func main() {

	var a, b int
	fmt.Scan(&a, &b)

	var c ???

	c = a == b
	
	if c {
		fmt.Println("uguali")
	} else {
		fmt.Println("diversi")
	}
}
```

Di che tipo deve essere la variabile `c` affinché la compilazione del codice non generi errori?