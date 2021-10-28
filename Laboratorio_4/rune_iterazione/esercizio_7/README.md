# Analisi lettere maiuscole/minuscole (1)

Scrivere un programma che legga da **standard input** una stringa senza spazi e, considerando **solamente** l’insieme delle lettere dell'alfabeto inglese, stampi
* il numero di lettere maiuscole;
* il numero di lettere minuscole;
* il numero di consonanti;
* il numero di vocali.

A tal fine definire le seguenti funzioni: 'èLetteraValida(l rune) bool', 'èMaiuscola(l rune) bool', 'èVocale(l rune) bool'.

*Suggerimento:* Le lettere dell'alfabeto inglese sono considerate nello standard US-ASCII (integrato nello standard Unicode). Per i caratteri considerati nello standard US-ASCII, il codice Unicode varia tra 0 e 127. In particolare, per le lettere dell’alfabeto inglese il codice Unicode varia negli intervalli:
* [65, 90] per le lettere MAIUSCOLE (con ‘A’=65 e ‘Z’=90)
* [97, 122] per le lettere minuscole (con ‘a’=97 e ‘z’=122)

Sia `c` una varibile di tipo `rune`, i seguenti blocchi di codici sono sintatticamente/semanticamente corretti:

```go
if (c >='A' && c<='Z') || (c>='a' && c<='z'){
	fmt.Println(string(c), "è una lettera dell’alfabeto inglese!")
}else{
	fmt.Println(string(c), "non è una lettera dell’alfabeto inglese!")
}
// Si noti che la funzione unicode.IsLetter(c) del package unicode potrebbe
// restituire 'true' anche per caratteri che corrispondono a lettere ma che 
// non appartenengono all’alfabeto inglese
``` 

##### Esempio d'esecuzione:

```text
$ go run analisi.go
Ciaoà
Maiuscole: 1
Minuscole: 3
Vocali: 3
Consonanti: 1

$ go run analisi.go
Certo!Sto,bene!ìììììì
Maiuscole: 2
Minuscole: 10
Vocali: 5
Consonanti: 7

$ go run analisi.go
aaAA
Maiuscole: 2
Minuscole: 2
Vocali: 4
Consonanti: 0
```
