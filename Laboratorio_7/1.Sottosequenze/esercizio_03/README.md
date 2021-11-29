# Esercizio 3

Scrivere un programma che:
* legga da **riga di comando** una stringa `s` costituita da cifre decimali;
* stampi a schermo tutte le sottosequenze (anche ripetute) della stringa `s`  nelle quali le cifre sono in ordine crescente (si considerino solamente sottosequenze di almeno 2 cifre).

Se la stringa `s` letta da **riga di comando** non è costituita solamente da cifre decimali, il programma termina senza stampare nulla.

Oltre alla funzione `main()`, deve essere definita ed utilizzata almeno la funzione `Sottostringhe(s string) []string`, che riceve in input 
un valore `string` nel parametro `s`, e restituisce un valore `[]string` che contiene tutte le sottosequenze desiderate

##### Esempio d'esecuzione:

```text
$ go run esercizio_3.go 123456
[12 123 1234 12345 123456 23 234 2345 23456 34 345 3456 45 456 56]

$ go run esercizio_3.go 654321
[]

$ go run esercizio_3.go 123121212
[12 123 23 12 12 12]

$ go run esercizio_3.go acc23

$ go run esercizio_3.go 01010101
[01 01 01 01]
```
