# Numeri perfetti

**Definizione**: Un numero naturale è perfetto se è uguale alla somma dei suoi divisori propri (per esempio, `6` è perfetto perché `6 = 1 + 2 + 3`).

Scrivere un programma che legga da **standard input** un numero intero `n` e stampi se `n` è perfetto oppure no.

Oltre alla funzione `main()`, il programma deve definire e utilizzare le seguenti funzioni:
* una funzione `SommaDivisoriPropri(n int) int` che riceve in input un valore intero nel parametro `n` e restituisce la somma dei divisori propri di `n`. Se `n` non ha divisori propri la funzione restituisce `0`;
* una funzione `ÈPerfetto(n int) bool` che riceve in input un valore intero nel parametro `n` e restituisce `true` se `n` è perfetto e `false` altrimenti; la funzione deve utilizzare la funzione `SommaDivisoriPropri()`.

##### Esempio d'esecuzione:

```text
$ go run numero_perfetto.go 
Inserisci un numero: 0
0 non è perfetto

$ go run numero_perfetto.go
Inserisci un numero: 1
1 non è perfetto

$ go run numero_perfetto.go
Inserisci un numero: 6
6 è perfetto

$ go run numero_perfetto.go
Inserisci un numero: 28
28 è perfetto
```

 
 