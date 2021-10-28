# Tabelline

Scrivere un programma che legga da **standard input** un numero intero e stampi la tabellina corrispondente solo se il numero inserito è compreso tra `1` e `9` (estremi inclusi).

Oltre alla funzione `main()`, il programma deve definire e utilizzare le seguenti funzioni:
* `Tabellina(numero int)` che riceve in input un valore intero nel parametro `numero` e, se `numero` è compreso tra `1` e `9` (estremi inclusi), stampa la tabellina associata, altrimenti non stampa nulla.

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

$ go run tabellina.go
Inserisci un numero: -1
```
