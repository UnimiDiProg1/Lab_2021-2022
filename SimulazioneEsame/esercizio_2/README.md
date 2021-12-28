# Esercizio 2

## Parte 1

Scrivere un programma che:
* legga da **standard input** una stringa `s`. Si assuma che la stringa non contenga caratteri di spaziatura.
* verifichi che la stringa sia composta di sole lettere minuscole nello standard di caratteri ASCII. Qualora non fosse così il programma deve terminare senza stampare nulla.
* stampi a schermo, tutte le sottosequenze della stringa `s`  nelle quali le lettere siano in ordine crescente (si considerino solamente sottosequenze di almeno 2 lettere).

Come mostrato nell'**Esempio d'esecuzione**, ciascuna sottosequenza deve essere stampata un'unica volta, riportando il numero di volte in cui la sottosequenza appare in `s`.

Oltre alla funzione `main()`, deve essere definita ed utilizzata almeno la funzione `Sottostringhe(s string) map[string]int`, che riceve in input 
un valore `string` nel parametro `s`, e restituisce un valore `map[string]int` in cui, per ogni sottosequenza di `s` di almeno 2 lettere, sia memorizzato il numero di volte in cui essa appare in `s`.

## Parte 2
Estendere la Parte 1 in modo tale che le sottosequenze vengano stampate in ordine lessicale crescente. A tale scopo, si utilizzi la libreria `sort`.

##### Esempio d'esecuzione:

```text
$ go run esercizio_2.go 
abcdef
output:
ab 1
abc 1
abcd 1
abcde 1
abcdef 1
bc 1
bcd 1
bcde 1
bcdef 1
cd 1
cde 1
cdef 1
de 1
def 1
ef 1

$ go run esercizio_2.go 
fedcba
output:

$ go run esercizio_2.go 
abcababab
output:
ab 4
abc 1
bc 1

$ go run esercizio_2.go 
abbèh

$ go run esercizio_2.go 
01010101

$ go run esercizio_2.go 
ABCD
```
