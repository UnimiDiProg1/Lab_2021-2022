# Operazioni

Create un programma che legge da **standard input** due numeri interi, che chiameremo `x` e `y`.
Letti i due numeri, il programma stampa:
* il maggiore tra `x` e `y`
* il minore tra `x` e `y`
* il risultato della somma tra `x` e `y`
* il risultato della differenza tra il maggiore e il minore 
* il risultato della divisione `x/y`
* il risultato del prodotto `x*y`
* il valore medio tra `x` e `y`
* il risultato di `x` elevato alla `y` (utilizzando sia un ciclo `for` sia la funzione `math.Pow`)

##### Esempio d'esecuzione:

```text
$ go run operazioni.go 
Inserisci due numeri interi: 
2 4
Maggiore: 4
Minore: 2
Somma: 6
Differenza: 2
Prodotto: 8
Divisione: 0.5
Valore medio: 3
Potenza (ciclo for): 16
Potenza (math.Pow): 16
```
