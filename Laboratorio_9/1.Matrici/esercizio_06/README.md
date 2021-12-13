# Coppie

Scrivere un programma che:
* legga da **riga di comando** un numero intero `n`;
* utilizzi le funzioni `CreaSliceBidimensionale(l int) [][]int` e `InizializzaSliceBidimensionale([][]int) ` descritte nell'Esercizio 2 (Laboratorio 9 - Array e slice III) per inizializzare una variabile `s` di tipo `[][]int` con lunghezza/capacità pari a `n` in cui ogni elemento `s[i]`, con `0 <= i < l`, è un valore di tipo `[]int` con lunghezza/capacità pari a `n`;
* stampi a video tutte le coppie di indici `(i, j)`, con `0 <= i < l` e `0 <= j < l`, tali che `s[i][j]` è uguale a `1`.

Oltre alle funzioni `main()`, `CreaSliceBidimensionale()`, e `InizializzaSliceBidimensionale()` devono essere definite ed utilizzate almeno le seguenti funzioni:
* una funzione `Coppie(s [][]int) (coppie [][]int)` che riceve in input un valore `[][]int` nel parametro `s` e restituisce il valore di tipo `[][]int` nella variabile `coppie` in cui sono memorizzate tutte le coppie di indici `(i, j)`, con `0 <= i < l` e `0 <= j < l`, tali che `s[i][j]` è uguale a `1` (`coppie[k]`, con `0 <= k < len(coppie)`, è un valore di tipo` []int` di lunghezza `2`).

##### Esempio d'esecuzione:

```text
$ go run coppie.go 4
[1 1]
[1 3]
[2 1]
[2 2]
[3 2]
[3 3]
```
