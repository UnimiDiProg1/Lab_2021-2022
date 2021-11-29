# Testo invertito

Scrivere un programma che legga da **standard input** un testo formato da un numero variabile di righe, terminando la lettura quando viene inserita da **standard input** una riga vuota (`""`), e lo ristampi dall’ultimo carattere al primo.

Scrivere un programma che: 
* legga da **standard input** un testo su più righe;
* termini la lettura quando viene inserita da **standard input** una riga vuota (`""`);
* ristampi il testo letto (riga vuota esclusa) dall’ultimo carattere al primo.

Oltre alla funzione `main()`, il programma deve definire e utilizzare le seguenti funzioni:
* una funzione `LeggiTesto() string` che legge da **standard input** un testo su più righe e terminato da una riga vuota (`""`), restituendo un valore `string` in cui è memorizzato il testo letto (riga vuota esclusa);
* una funzione `InvertiStringa(s string) string` che riceve in input un valore `string` nel parametro `s` e ne inverte l'ordine dei caratteri, ovvero restituisce un valore `string` in cui il primo carattere è l'ultimo che definisce `s`, il secondo carattere è il penultimo che definisce `s`, ...  e l'ultimo carattere è il primo che definisce `s`.

##### Esempio d'esecuzione:

```text
$ go run stringainvertita.go
Inserisci un testo su più righe (termina con riga vuota):
Testo di prova
disposto su due righe

Testo invertito:
ehgir eud us otsopsid
avorp id otseT
```
