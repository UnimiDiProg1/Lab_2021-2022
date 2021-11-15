# Arrotondamento

Scrivere un programma che legga da **standard input** un numero reale e ne stampi il valore **arrotondato** alla seconda cifra decimale.

*Suggerimento:* Il valore arrotondato alla seconda cifra decimale di un numero reale può essere ottenuto effettuando le seguenti operazioni:

1. Moltiplicare il numero reale per 100
2. Sommare 0.5 al valore ottenuto al passo 1) 
3. Convertire il valore ottenuto al passo 2) in un valore di tipo `int` 
4. Riconvertire il valore ottenuto al passo 3) in un valore di tipo `float64` 
5. Dividere il valore ottenuto al passo 4) per 100

oppure:

1. Moltiplicare il numero reale per 100
2. Utilizzare la funzione `math.Round` del package `math` per arrontondare all'intero valore ottenuto al passo 1) (si utilizzi il comando `go doc math.Round` per capire come utilizzare la funzione)
3. Dividere il valore ottenuto al passo 2) per 100

##### Esempio d'esecuzione:

```text
$ go run arrotondamento.go 
Inserisci il valore da arrotondare: 10.34762
Valore arrotondato = 10.35

$ go run arrotondamento.go
Inserisci il valore da arrotondare: 8.32467
Valore arrotondato = 8.32
``` 

