# Trasformazione lettere maiuscole/minuscole

Scrivere un programma che legga da **standard input** una stringa e, considerando l’insieme delle lettere dell'alfabeto inglese, ristampi a video la stringa due volte: la prima volta in maiuscolo e la seconda volta in minuscolo.

A tal fine definire due funzioni: 'inMaiuscolo(carattere rune) rune' e 'inMinuscolo(carattere rune) rune'

 *Suggerimento:* Sia `c` una varibile di tipo `rune`, i seguenti blocchi di codici sono sintatticamente/semanticamente corretti:
 ```go
 if (c >='A' && c <= 'Z') {
 	fmt.Println("L’equivalente lettera minuscola è:", string('a' + (c - 'A')))
 }
 // ... oppure, utilizzando il package "unicode"...
 if (c >='A' && c <= 'Z') {
 	fmt.Println("L’equivalente lettera minuscola è:", string(unicode.ToLower(c)))
 }
 ``` 

##### Esempio d'esecuzione:

```text
$ go run trasforma.go 
TestoDiProva!!!
Testo maiuscolo: TESTODIPROVA!!!
Testo minuscolo: testodiprova!!!

$ go run trasforma.go
Testo_Di_PrOvA....
Testo maiuscolo: TESTO_DI_PROVA....
Testo minuscolo: testo_di_prova....

```
