# Trova l'errore

Questo programma dovrebbe stampare 4 asterischi ma non funziona correttamente. Qual è l'errore? 

```go
package main

import "fmt"

func main() {
	var n int = 2

	var i, j int
	for i = 0; i < n; i++ {
    	for j = 0; j < n; j-- {
    		fmt.Print("*")
    	}
    }
    fmt.Println()
}
```
