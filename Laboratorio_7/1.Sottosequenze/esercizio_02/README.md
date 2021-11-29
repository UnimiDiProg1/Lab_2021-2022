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
