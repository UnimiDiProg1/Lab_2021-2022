# Qual è l'output?

Supponendo che l'utente inserisca da **riga di comando** il valore `tabellina.txt`, cosa dovrebbe produrre in output il seguente programma?

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create(os.Args[1])
	if err != nil {
		fmt.Printf("Error while creating the file! %v\n", err)
		return
	}
	defer f.Close()

	for i:=1; i<=10; i++ {
		risultato := i * 10
		_, err = fmt.Fprintln(f, "10 x ", i, risultato)
        // ... equivalente a:
        // _, err = fmt.Fprintf(f, "10 x %d = %d\n", i, risultato)
		if err != nil {
			fmt.Printf("Error while writing the file! %v\n", err)
			return
		}
	}
}
```
