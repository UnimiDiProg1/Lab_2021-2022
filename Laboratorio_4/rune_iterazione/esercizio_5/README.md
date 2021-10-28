# Qual è l'output?

Qual è l'output del seguente programma?

```go
package main

import "fmt"

func main() {

	s := "Ciao, come va?"
    // s è interamente definita da caratteri considerati nello standard US-ASCII

	for i := 0; i<len(s); i++ {
		fmt.Print(string(s[i]))
    }
    fmt.Println()

	s = "Ciao, come è andata?"
    // s non è interamente definita da caratteri considerati nello standard US-ASCII

	for i := 0; i<len(s); i++ {
		fmt.Print(string(s[i]))
    }
    fmt.Println()

}
```

Osservazioni:
* Data una stringa `s`, `s[i]` è il byte in posizione `i` nella sequenza di byte che rappresenta `s`. In generale, `s[i]` **non** codifica un carattere.
* In generale, per esaminare in sequenza i caratteri che definiscono una stringa si deve utilizzare il costrutto `for range`.
