# Garibaldi

Scrivere un programma che: 
* legga da **standard input** un testo su più righe;
* termini la lettura quando viene inserita da **standard input** una riga vuota (`""`);
* ristampi il testo letto (riga vuota esclusa) sostituendo tutte le vocali con delle `u`.

Oltre alla funzione `main()`, il programma deve definire e utilizzare le seguenti funzioni:
* una funzione `LeggiTesto() string` che legge da **standard input** un testo su più righe e terminato da una riga vuota (`""`), restituendo un valore `string` in cui è memorizzato il testo letto (riga vuota esclusa);
* una funzione `TrasformaCarattere(c rune) rune` che riceve in input un valore `rune` nel parametro `c` e restituisce un valore `rune` uguale a `u` se `c` è una vocale, e uguale a `c` altrimenti.
* una funzione `Garibaldi(s string) string` che riceve in input un valore `string` nel parametro `s` e restituisce un valore `string` in cui ogni carattere è trasformato utilizzando la funzione `TrasformaCarattere()`.

##### Esempio d'esecuzione:

```text
$ go run garibaldi.go
Inserisci un testo su più righe (termina con riga vuota):
Garibaldi fu ferito
fu ferito in una gamba,
Garibaldi che comanda
che comanda i bersaglier

Risultato trasformazione:
Gurubuldu fu furutu
fu furutu un unu gumbu,
Gurubuldu chu cumundu
chu cumundu u bursugluur
```
