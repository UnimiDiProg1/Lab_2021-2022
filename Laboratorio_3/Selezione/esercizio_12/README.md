# Conversioni

Scrivere un unico programma che: 
- legga da **standard input** un valore intero che specifica il tipo di conversione da effettuare:  
1: secondi (inseriti dall’utente) in ore  
2: secondi inseriti dall’utente in minuti  
3: minuti inseriti dall’utente in ore  
4: minuti inseriti dall’utente in secondi  
5: ore inserite dall’utente in secondi  
6: ore inserite dall’utente in minuti  
7: minuti inseriti dall’utente in giorni e ore  
8: minuti inseriti dall’utente in anni e giorni

    gestendo l’insertimento di un valore di scelta non compreso tra 1 e 8;

- legga da **standard input** un valore reale da convertire;

- stampi a video il valore convertito.

##### Esempio d'esecuzione:

```text
$ go run conversioni.go 
Scegli la conversione:
1) secondi -> ore
2) secondi -> minuti
3) minuti -> ore
4) minuti -> secondi
5) ore -> secondi
6) ore -> minuti
7) minuti -> giorni e ore
8) minuti -> anni e giorni
: 8
Inserisci il valore da convertire: 7200
7200 minuti corrispondono a 0 anni e 5 giorni

$ go run conversioni.go
Scegli la conversione:
1) secondi -> ore
2) secondi -> minuti
3) minuti -> ore
4) minuti -> secondi
5) ore -> secondi
6) ore -> minuti
7) minuti -> giorni e ore
8) minuti -> anni e giorni
: 1
Inserisci il valore da convertire: 3618
3618 secondi corrispondono a 1.005 ore

$ go run conversioni.go
Scegli la conversione:
1) secondi -> ore
2) secondi -> minuti
3) minuti -> ore
4) minuti -> secondi
5) ore -> secondi
6) ore -> minuti
7) minuti -> giorni e ore
8) minuti -> anni e giorni
: 9
Scelta errata
```
