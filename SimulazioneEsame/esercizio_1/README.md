# Esercizio 1

## Parte 1

Scrivere un programma che legga da **riga di comando** una stringa rappresentante una funzione quadratica nella forma ax^2+bx+c=0, con a,b,c interi e diversi da 0.

**Esempio:**

`5x^2+1x+2=0`

Il programma deve stampare su standard output quante e quali sono le soluzioni reali di questa equazione. Ogni soluzione deve essere approssimata usando 3 cifre decimali. Le soluzioni si possono trovare applicando la seguente formula:

<img src="https://ewserver.di.unimi.it/marble/equation.jpg" alt="x=\frac{-b\pm\sqrt{b^2-4ac}}{2a}" width="150"/>

Ne risulta che il numero di soluzioni reali totali di un’equazione quadratica possono essere zero, se il discriminante, ovvero la parte sotto la radice quadrata, è negativo; una, se il discriminante è nullo; oppure due altrimenti.
Si ricorda che la funzione `sqrt` per calcolare la radice quadrata è disponibile nel package `math`.
È possibile utilizzare la funzione `split` del package `strings` per suddividere una stringa in parti separate da una sottostringa specificata.

## Parte 2

Estendere la Parte 1 in modo tale da leggere da **riga di comando** anche due valori reali `soglia` ed `epsilon` e stampare quali soluzioni siano maggiori del valore `soglia` di almeno `epsilon`

##### Esempio d'esecuzione:
```text
$ go run esercizio_1.go 4x^2+10x+1=0 -1 0.01
Esistono due soluzioni reali: -0.104 e -2.396
La soluzione -0.104 è maggiore della soglia

$ go run esercizio_1.go 4x^2-10x+1=0 2.3 0.05
Esistono due soluzioni reali: 2.396 e 0.104
La soluzione 2.396 è maggiore della soglia

$ go run esercizio_1.go 1x^2+10x+25=0 -6 1
Esiste un’unica soluzione reale: -5.000

$ go run esercizio_1.go 5x^2+1x+2=0
Non ci sono soluzioni reali
```
