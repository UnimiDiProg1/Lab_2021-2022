# Qual è l'output?

Supponendo che l'utente inserisca da **standard input** il valore `10`, cosa dovrebbe produrre in output il seguente programma? E se invece inserisse `4`?

```go
package main

import "fmt"

func main() {

	var voto int

	fmt.Scan(&voto)

	switch voto {
	default:
		fmt.Println("Insufficiente!")
	case 10:
		fallthrough
	case 9:
		fmt.Println("Ottimo!")
	case 8:
		fmt.Println("Distinto!")
	case 7:
		fmt.Println("Buono!")
	case 6:
		fmt.Println("Sufficiente!")
	}
}
```
 # Qual è l'output?

Supponendo che l'utente inserisca da **standard input** il valore `9`, cosa dovrebbe produrre in output il seguente programma? E se invece inserisse `13`?

```go
package main

import "fmt"

func main() {

	var n int

	fmt.Scan(&n)

	var somma int

	for i:=0; i<=n; i++ {

		switch i%2 {
		case 0:
			fmt.Println(i, "pari!")
		case 1:
			continue
		}

		somma += i
	}

	fmt.Println("Somma =", somma)
}
```
 # Qual è l'output?

Supponendo che l'utente inserisca da **standard input** il valore `9`, cosa dovrebbe produrre in output il seguente programma? E se invece inserisse `13`?

```go
package main

import "fmt"

func main() {

	var n int

	fmt.Scan(&n)

	var somma int

	for i:=0; i<=n; i++ {

		switch i%2 {
		case 0:
			fmt.Println(i, "pari!")
		case 1:
			break
		}

		somma += i
	}

	fmt.Println("Somma =", somma)
}
```
 # Qual è l'output?

Supponendo che l'utente inserisca da **standard input** il valore `3`, cosa dovrebbe produrre in output il seguente programma? E se invece inserisse `100`?

```go
package main

import "fmt"

func main() {

	var n int

	fmt.Scan(&n)

	accumulatore := 1

	for i:=1; i<=n; i++ {

		switch {
		case i<5:
			accumulatore *= i
		case i<10:
			accumulatore += i
		}

	}

	fmt.Println("Accumulatore =", accumulatore)
}
```
 # Analisi lettere maiuscole/minuscole (2)

Scrivere un programma che legga da **standard input** una stringa senza spazi e, considerando l’insieme delle lettere dell'alfabeto inglese, stampi
* il numero di vocali maiuscole;
* il numero di vocali minuscole;
* il numero di consonanti maiuscole;
* il numero di consonanti minuscole.

A tal fine definire le seguenti funzioni: 'èLetteraValida(l rune) bool', 'èMaiuscola(l rune) bool', 'èVocale(l rune) bool'.

Modificare il programma (già svolto la scorsa volta) per
* Utilizzare il costrutto switch per il controllo delle vocali nella funzione 'èVocale'
* Utilizzare le funzioni 'unicode.IsLetter' e 'unicode.IsUpper' del package 'unicode' al posto di 'èLetteraValida' e 'èMaiuscola' rispettivamente.

##### Esempio d'esecuzione:

```text
$ go run analisi.go   
Ciao
Vocali maiuscole: 0
Consonanti maiuscole: 1
Vocali minuscole: 3
Consonanti minuscole: 0

$ go run analisi.go        
Certo!Sto,bene
Vocali maiuscole: 0
Consonanti maiuscole: 2
Vocali minuscole: 5
Consonanti minuscole: 5

$ go run analisi.go
aaAA
Vocali maiuscole: 2
Consonanti maiuscole: 0
Vocali minuscole: 2
Consonanti minuscole: 0
```
# Calcolatrice

Scrivere un programma che legga da **standard input** un numero reale, `a`, un simbolo di operazione aritmetica (`+`, `-`, `*` o `/`), ed un secondo numero reale `b`.
Il programma deve calcolare e stampare a video il risultato dell'operazione specificata tra i due numeri. Se l'operatore non è riconosciuto il programma invece deve stampare `Operatore non riconosciuto`.
Si utilizzi il costrutto switch per selezionare l'operazione in base al simbolo inserito.

##### Esempio d'esecuzione:

```text
$ go run calcolatrice.go
inserisci un'operazione aritmetica
3.1 + 6.0
risultato: 9.1

$ go run calcolatrice.go
inserisci un'operazione aritmetica
4 - 2
risultato: 2

$ go run calcolatrice.go
inserisci un'operazione aritmetica
10 * 4.5
risultato: 45

$ go run calcolatrice.go
inserisci un'operazione aritmetica
10 / 2
risultato: 5

$ go run calcolatrice.go
inserisci un'operazione aritmetica
10 % 2
Operatore non riconosciuto
```

# Mesi

Scrivere un programma che legga da **standard input** il nome di un mese (in minuscolo).
Il programma deve stampare a video il numero di giorni di quel mese. Si assume che l'anno non sia bisestile.

In particolare, si scriva una funzione `numeroDiGiorni` che, dato in input il nome del mese (variabile di tipo stringa) restituisca il numero di giorni del mese (variabile di tipo intera).
La funzione deve usare il costrutto switch case.


##### Esempio d'esecuzione:

```text
$ go run mesi.go
Inserire mese: ottobre
Il numero di giorni di ottobre è 31

$ go run calcolatrice.go
Inserire mese: aprile
Il numero di giorni di aprile è 30

$ go run calcolatrice.go
Inserire mese: catalogna
Il nome del mese non è corretto.
```
