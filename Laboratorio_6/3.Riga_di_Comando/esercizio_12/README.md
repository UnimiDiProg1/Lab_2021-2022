# Divisori comuni

**Definizione**: I divisori propri di un numero naturale (un numero intero positivo) sono tutti i suoi divisori, tranne il numero stesso.
 
**Definizione**: Un numero naturale (un numero intero positivo) è perfetto se è uguale alla somma dei suoi divisori propri (per esempio, 6 è perfetto perché 6 = 1 + 2 + 3 ).
 
Scrivere un programma che:
* legga da **riga di comando** tre numeri interi positivi, rispettivamente `N`, `DIVISORIMIN`, e `DIVISORIMAX`;
* stampi a video tutte le coppie di interi positivi `a` e `b`, con `a <= N` e `b <= N`, tali che:
    1. `a` e `b` abbiano in comune un numero di divisori propri maggiore o uguale a `DIVISORIMIN` e minore o uguale a `DIVISORIMAX`;
    2. almeno uno tra `a` e `b` sia un numero perfetto.
     
Oltre alla funzione `main()`, devono essere definite ed utilizzate almeno le seguenti funzioni:

* una funzione `DivisoriPropri(n int) []int` che riceve in input un valore `int` nel parametro `n` e restituisce un valore `[]int` in cui sono memorizzati tutti i divisori propri di `n`.

Si assuma che:
* i valori letti da **riga di comando** siano specificati nel formato corretto;
* `N > MAX > MIN > 0`.


##### Esempio d'esecuzione:

```text
$ go run divisori.go 50 3 4
6 6
6 12
6 18
6 24
6 30
6 36
6 42
6 48
8 28
12 28
14 28
16 28
20 28
24 28
28 32
28 36
28 40
28 42
28 44
28 48

$ go run divisori.go 20 1 2
2 6
3 6
4 6
5 6
6 7
6 8
6 9
6 10
6 11
6 13
6 14
6 15
6 16
6 17
6 19
6 20

$ go run divisori.go 20 2 3
4 6
6 6
6 8
6 9
6 10
6 12
6 14
6 15
6 16
6 18
6 20
```
