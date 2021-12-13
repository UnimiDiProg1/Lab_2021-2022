# Coppie

Scrivere un programma che:
* legga da **riga di comando** un numero intero `n`;
* utilizzi le funzioni `CreaSliceBidimensionale(l int) [][]int` e `AssegnaSliceBidimensionale(s [][]int) ` descritte nell'Esercizio 4 per inizializzare una variabile `s` di tipo `[][]int` con lunghezza/capacità pari a `n` in cui ogni elemento `s[i]` è un valore di tipo `[]int` con lunghezza/capacità pari a `n`;
* stampi a video tutte le coppie di indici `(i, j)`, tali che `s[i][j]` è uguale a `1`.

Oltre alle funzioni `main()`, `CreaSliceBidimensionale()`, e `AssegnaSliceBidimensionale()` devono essere definite ed utilizzate almeno le seguenti funzioni:
* una funzione `Coppie(s [][]int) (coppie [][]int)` che riceve in input un valore `[][]int` nel parametro `s` e restituisce il valore di tipo `[][]int` nella variabile `coppie` in cui sono memorizzate tutte le coppie di indici `(i, j)`, tali che `s[i][j]` è uguale a `1`.

##### Esempio d'esecuzione:

```text
$ go run coppie.go 3
la slice originale è:
        0       1       1
        1       1       0
        1       1       0
gli indirizzi degli elementi uguali a 1 sono:
[0 1]
[0 2]
[1 0]
[1 1]
[2 0]
[2 1]
```
