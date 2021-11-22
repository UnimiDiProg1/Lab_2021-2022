# Esercizio 1

Scrivere un programma che legga da **riga di comando** una stringa composta di caratteri unicode e stampi a schermo tutte le sottostringhe palindrome (che siano uguali lette da destra e da sinistra) della stringa.

##### Esempio d'esecuzione
```text
$ go run esercizio_1.go sottotono
[otto tt tot oto ono]

$ go run esercizio_1.go banana
[ana anana nan ana]

$ go run esercizio_1.go ereggere
[ere ereggere regger egge gg ere]

$ go run esercizio_1.go Vaðlaheiðarvegavinnuverkfærageymsluskúrslyklakippuhringurinn
[nn pp nn]

$ go run esercizio_1.go donaudampfschifffahrtselektrizitätenhauptbetriebswerkbauunterbeamtengesellschaft
[ff fff ff ele izi tät uu ese ll]
```
# Esercizio 2

Uno dei compiti più importanti di un compilatore è quello di controllare se le parentesi sono ben bilanciate. Ad ogni parentesi aperta deve corrispondere una parentesi chiusa, e le coppie di parentesi devono essere innestate propriamente.

Un esempio di parentesi tonde ben bilanciate è il seguente: 
`(())()`

Un esempio di parentesi tonde *non* ben bilanciate è il seguente:
`())(`

Notare che in quest’ultimo esempio, anche se ad ogni parentesi aperta corrisponde una parentesi chiusa, le coppie di parentesi non sono propriamente innestate (viene chiusa una parentesi di troppo ed una parentesi rimane aperta senza mai essere chiusa).

## Parte 1
Leggere da **riga di comando** una stringa. Se la stringa contiene caratteri diversi da parentesi tonda aperta `(` e parentesi tonda chiusa `)` (che ricordiamo sono caratteri ASCII) terminare l’esecuzione del programma, altrimenti stampare `bilanciata` se la stringa ha parentesi bilanciate o `non bilanciata` altrimenti.
A tal fine implementare e usare una funzione `isBalanced(sequence string) bool` che riceve in input una stringa `sequence` composta di sole parentesi (aperte e chiuse) e che restituisce il valore booleano true se la stringa sequence è ben bilanciata, o false altrimenti.

* Nota: poichè le parentesi sono caratteri speciali nell'esecuzione dei programmi da riga di comando, per inserirle bisogna o fare l'escape di ciascuna oppure inserire la stringa tra virgolette.

## Parte 2
Stampare tutte le sottostringhe ben bilanciate.

##### Esempio d'esecuzione:
```text
$ go run esercizio_2.go "(())()"
bilanciata
---
Sottosequenze bilanciate:
1) (())
2) (())()
3) ()
4) ()

$ go run esercizio_2.go "(()"
non bilanciata
---
Sottosequenze bilanciate:
1) ()

$ go run esercizio_2.go "(()())"
bilanciata
---
Sottosequenze bilanciate:
1) (()())
2) ()
3) ()()
4) ()

non bilanciata
---
Sottosequenze bilanciate:
1) ()

$ go run esercizio_2.go "())("
non bilanciata
---
Sottosequenze bilanciate:
1) ()

go run esercizio_2.go "))(("
non bilanciata
---
Sottosequenze bilanciate:
Nessuna.

go run esercizio_2.go []
L'input fornito non aveva esclusivamente parentesi tonde.
```
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
