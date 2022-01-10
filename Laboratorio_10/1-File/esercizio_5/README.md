# Potenze

Scrivere un programma che legga da **riga di comando** due numeri interi `b` ed `e` e stampi su un nuovo file di testo `potenze.txt` le potenze del numero `b` fino a `b`<sup>`e`</sup>.

*Esempio*: Se `b=2` ed `e=5`, il programma deve stampare sul file `potenze.txt` i numeri `1 2 4 8 16 32`.

##### Esempio d'esecuzione:

```text
$ go run potenze.go 2 5

$ cat potenze.txt      
1 2 4 8 16 32

$ go run potenze.go 4 3

$ cat potenze.txt      
1 4 16 64
```