# Ripetizioni


Scrivere un programma che legga da **riga di comando** una stringa di caratteri `file.txt`.

`file.txt` è il nome di un file di testo memorizzato nella stessa directory in cui è memorizzato il programma.

Il programma deve stampare su un nuovo file di testo `statistiche.txt`: 
1. Il numero di parole distinte presenti nel file `file.txt` (una parola è una stringa interamente definita da caratteri il cui codice Unicode, se passato come argomento alla funzione `func IsSpace(r rune) bool` del package `unicode`, fa restituire `false` alla funzione).
2. La lista di parole distinte presenti nel file `file.txt`, riportando per ogni parola ripetuta il relativo numero di occorrenze (cfr. **Esecuzione d'esecuzione**).

##### Esempio d'esecuzione:

```text
$ cat testo.txt 
Lontano, nei dimenticati spazi non segnati nelle carte geografiche 
dell'estremo limite della Spirale Ovest della Galassia, c'è un piccolo 
e insignificante sole giallo.
A orbitare intorno a esso, alla distanza di centoquarantanove milioni 
di chilometri, c'è un piccolo, trascurabilissimo pianeta azzurro–verde, 
le cui forme di vita, discendenti dalle scimmie, sono così incredibilmente
primitive che credono ancora che gli orologi da polso digitali siano
un'ottima invenzione.

Guida galattica per gli autostoppisti

$ go run ripetizioni.go testo.txt

$ more statistiche.txt
Totale parole distinte: 62
Occorrenze:
LONTANO,: 1
DELLA: 2
DALLE: 1
CHE: 2
SEGNATI: 1
DELL'ESTREMO: 1
A: 2
CHILOMETRI,: 1
PIANETA: 1
UN'OTTIMA: 1
COSÌ: 1
--More--(21%)
``` 