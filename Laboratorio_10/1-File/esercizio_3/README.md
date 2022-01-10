# Qual è l'output?

Supponendo che:
* l'utente inserisca da **riga di comando** i valori `testo.in testo.out` ,
* il file `testo.in` contenga il seguente testo:
```text
La nebbia a gl'irti colli
piovigginando sale,
e sotto il maestrale
urla e biancheggia il mar;
```
* il file `testo.in` sia memorizzato nella stessa directory in cui è memorizzato il programma di seguito riportato,

cosa dovrebbe produrre in output il programma? 

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		fIn, fOut *os.File
		err error
	)

	fIn, err = os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("Error while opening the file! %v\n", err)
		return
	}
	defer fIn.Close()
	fOut, err = os.Create(os.Args[2])
	if err != nil {
		fmt.Printf("Error while creating the file! %v\n", err)
		return
	}
	defer fOut.Close()

	scanner := bufio.NewScanner(fIn)
	counter := 0
	for scanner.Scan() {
		line := scanner.Text()
		counter++
		_, err = fmt.Fprintf(fOut,"LINE %3d: <%s>\n", counter, line)
		if err != nil {
			fmt.Printf("Error while writing the file! %v\n", err)
			return
		}
	}
	if err = scanner.Err(); err != nil {
		fmt.Printf("Error while reading the file! %v\n", err)
		return
	}
}
```
