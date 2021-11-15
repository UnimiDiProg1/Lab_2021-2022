# Troncamento

Scrivere un programma che legga da **standard input** un numero reale e ne stampi il valore **troncato** alla seconda cifra decimale.

*Suggerimento:* Il valore troncato alla seconda cifra decimale di un numero reale può essere ottenuto effettuando le seguenti operazioni:  
1. Moltiplicare il numero reale per 100   
2. Convertire il valore ottenuto al passo 1) in un valore di tipo `int`   
3. Riconvertire il valore ottenuto al passo 2) in un valore di tipo `float64`   
4. Dividere il valore ottenuto al passo 3) per 100  

oppure:

1. Moltiplicare il numero reale per 100  
2. Utilizzare la funzione `math.Trunc` del package `math` per troncare all'intero valore ottenuto al passo 1) (si utilizzi il comando `go doc math.Trunc` per capire come utilizzare la funzione)  
3. Dividere il valore ottenuto al passo 2) per 100 

##### Esempio d'esecuzione:

```text
$ go run troncamento.go 
Inserisci il valore da troncare: 10.34762
Valore troncato = 10.34

$ go run troncamento.go        
Inserisci il valore da troncare: 8.34267
Valore troncato = 8.34
``` 

