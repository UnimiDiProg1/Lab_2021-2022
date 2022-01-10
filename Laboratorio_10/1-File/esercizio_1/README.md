# Qual è l'output?

Supponendo che:
* l'utente inserisca da **riga di comando** il valore `punti.txt`,
* il file `punti.txt` contenga il seguente testo:
```text
A 10.5 20
B 15 30
C 12.5 25.6
```
* il file `punti.txt` sia memorizzato nella stessa directory in cui è memorizzato il programma di seguito riportato,

cosa dovrebbe produrre in output il programma? 

```go
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("Error while opening the file! %v\n", err)
		return
	}
	defer f.Close()

	for {
		var nome string
		var x, y float64
		_, err = fmt.Fscan(f, &nome, &x, &y)
        // ... equivalente a:
        // _, err = fmt.Fscanf(f, "%s%f%f", &nome, &x, &y)
        // ... oppure a:
        // _, err = fmt.Fscanf(f, "%s %f %f", &nome, &x, &y)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Error while reading the file! %v\n", err)
			return
		}
		fmt.Printf("Punto %s = (%v, %v)\n", nome, x, y)
	}
}
```
