# Qual è l'output?

Analizziamo l'output del seguente programma.

```go
package main

import "fmt"

func main() {
	mappa := make(map[string]int)

	mappa["A"] = 10
	mappa["B"] -= 5
	mappa["D"] = mappa["E"] + 5
	if mappa["F"] == 0 {
	
		fmt.Printf("F è presente con valore %d\n", mappa["F"])
	
	} else {
	
		fmt.Print("F non è presente\n")
	
	}

	fmt.Println("Elementi in mappa:")
	
	for k := range mappa {

		fmt.Printf("Chiave: %s - Valore: %d\n", k, mappa[k])
		
	}
}
```
