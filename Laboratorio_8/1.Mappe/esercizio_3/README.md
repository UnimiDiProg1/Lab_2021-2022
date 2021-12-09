# Qual è l'output?

Analizziamo l'output del seguente programma.

```go
package main

import (
	"fmt"
)

var (
	nominativi = map[string]string{"023314944": "Mario Rossi", 
		"024158685": "Carlo Bianchi", "026424971": "Giuseppe Verdi", 
		"0269001634": "Carlo Bianchi", "026691369": "Mario Rossi",
		"0248704925": "Carlo Bianchi", "023554756": "Giuseppe Verdi"}
)

func main() {
	numeriTelefonici := make(map[string]string)
	for k, v := range nominativi {
		numeriTelefonici[v] = k
	}
	for k, v := range numeriTelefonici {
		fmt.Printf("Nominativo: %v\nNumero telefonico: %v\n\n", k, v)
	}
}
```