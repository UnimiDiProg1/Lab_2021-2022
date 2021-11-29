# Statistiche testo

Scrivere un programma che: 
* legga da **standard input** un testo su più righe;
* termini la lettura quando, premendo la combinazione di tasti `Ctrl+D`, viene inserito da **standard input** l'indicatore End-Of-File (EOF);
* stampi a video le seguenti statistiche relative al testo letto:
  1. il numero di linee
  1. il numero di parole presenti nel testo;
  2. la lunghezza media delle parole presenti nel testo.

**Nota:** una parola è una sequenza di caratteri consecutivi rappresentanti delle lettere. I caratteri di spaziatura intervallano parole diverse.

Oltre alla funzione `main()`, il programma deve definire e utilizzare le seguenti funzioni:
* una funzione `LeggiTesto() string` che legge da **standard input** un testo su più righe terminato dall'indicatore EOF (`CTRL+D`), restituendo un valore `string` in cui è memorizzato il testo letto;
* una funzione `StatisticheParole(s string) (int, int, int)` che riceve in input un valore `string` nel parametro `s` e restituisce tre valori `int` pari rispettivamente al numero di linee presenti in `s`, il numero  parole presenti in `s` e alla loro lunghezza totale.

*Suggerimento:* per dividere una stringa in una slice di sottostringhe in base ad un carattere separatore, si utilizzi la funzione `strings.Split()` del package `strings`. Questo può essere utile per dividere una stringa isolando le linee (carattere separatore '`\n`') o per ottenere le parole separate da uno spazio (carattere separatore '` `'). Utilizzare `go doc strings.Split` per scoprire il suo funzionamento.

##### Esempio d'esecuzione:

```text
$ go run statistiche.go    
Inserisci un testo su più righe (termina con Ctrl+D):
testo di prova
su cui provare le statistiche
Statistiche:
Numero linee: 2
Numero parole: 8
Lunghezza media: 4.625
```
