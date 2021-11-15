# Troncamento/Arrotondamento generalizzati

Scrivere una versione generalizzata dei programmi **Troncamento** e **Arrotondamento** che legga da **standard input** un intero `n` oltre al numero reale. L'intero `n` specifica che il troncamento e l'arrotondamento devono avvenire alla `n`-esima cifra decimale.

##### Esempio d'esecuzione:

```text
$ go run generalizzazione.go
Inserisci il valore: 10.34762
Inserisci il numero di cifre dopo la virgola: 4
Valore troncato = 10.3476
Valore arrotondato = 10.3476

$ go run generalizzazione.go
Inserisci il valore:  10.34762
Inserisci il numero di cifre dopo la virgola: 3
Valore troncato = 10.347
Valore arrotondato = 10.348
``` 

