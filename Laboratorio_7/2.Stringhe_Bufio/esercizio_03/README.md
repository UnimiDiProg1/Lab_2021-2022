# Testo concatenato

Scrivere un programma che: 
* legga da **standard input** un testo su più righe;
* termini la lettura quando, premendo la combinazione di tasti `Ctrl+D`, viene inserito l'indicatore End-Of-File (EOF);
* ristampi il testo letto.

**Nota:** la stampa del testo deve essere effettuata solamente dopo aver letto per intero il testo.  

Oltre alla funzione `main()`, il programma deve definire e utilizzare le seguenti funzioni:
* una funzione `LeggiTesto() string` che legge da **standard input** un testo su più righe e terminato dall'indicatore EOF, restituendo un valore `string` in cui è memorizzato il testo letto.

La funzione `LeggiTesto() string` deve utilizzare uno scanner di standard input inizializzato mediante la funzione `scanner := bufio.NewScanner(os.Stdin)`, resa disponibile dal package `bufio`.
Per leggere una riga di testo inserita da standard input utilizzare il metodo `scanner.Scan()` reso disponibile sullo scanner appena creato.
Per accedere alla riga di testo appena letta, utilizzare il metodo `scanner.Text()` reso disponibile sullo scanner appena creato.

##### Esempio d'esecuzione:

```text
$ go run ristampa.go 
Inserisci testo (termina con CTRL+D):
Testo di prova
disposto su più righe

Alcune possono essere anche righe vuote
Testo letto:
Testo di prova
disposto su più righe

Alcune possono essere anche righe vuote
```
