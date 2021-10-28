# Tabelline 2

*Estendete il programma che risolve l'esercizio **Tabelline** in modo tale che soddisfi la richiesta seguente.* 

Scrivere un programma che legga da **standard input** un numero intero e stampi la tabellina corrispondente solo se il numero inserito è compreso tra `1` e `9` (estremi inclusi). Se il numero inserito non è valido (inferiore di `1` o superiore a `9`), il programma deve stampare `Numero non valido.`.

Oltre alla funzione `main()`, il programma deve definire e utilizzare le seguenti funzioni:
* `Tabellina(numero int) bool` che riceve in input un valore intero nel parametro `numero`. Se `numero` è compreso tra `1` e `9` (estremi inclusi), la funzione stampa la tabellina associata e restituisce un valore logico `true`. Altrimenti, la funzione non stampa nulla e restituisce un valore logico `false`.

##### Esempio d'esecuzione:

```text
$ go run tabellina.go
Inserisci un numero: 4 
Tabellina del 4: 0 4 8 12 16 20 24 28 32 36 40

$ go run tabellina.go
Inserisci un numero: 6
Tabellina del 6: 0 6 12 18 24 30 36 42 48 54 60

$ go run tabellina.go
Inserisci un numero: 13
Numero non valido.

$ go run tabellina.go
Inserisci un numero: -1
Numero non valido.
```
