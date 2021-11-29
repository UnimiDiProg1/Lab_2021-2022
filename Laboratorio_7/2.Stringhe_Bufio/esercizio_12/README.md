# Date (1)

Scrivere un programma che:
* legga da **standard input** un testo su più righe;
* termini la lettura quando viene inserita da standard input una riga vuota (`""`).

Ogni riga di testo è una una stringa senza spazi che codifica una data in uno dei seguenti possibili formati:

1. aaaa/m/g
2. aaaa/m/gg
3. aaaa/mm/g
4. aaaa/mm/gg

Una volta terminata la fase di lettura, il programma deve stampare la codifica nel formato aaaa-mm-gg delle date lette.

Oltre alla funzione `main()` devono essere definite ed utilizzate almeno le seguenti funzioni:
* una funzione `LeggiTesto() []string` che legge da **standard input** un testo su più righe e terminato da una riga vuota (`""`), restituendo un valore `[]string` in cui sono memorizzate le righe del testo letto;
* una funzione `Formatta(s string) string` che riceve in input un valore `string` nel parametro `s` in cui è codificata una data nel formato **aaaa/m/g**, **aaaa/m/gg**, **aaaa/mm/g** o **aaaa/mm/gg**, e restituisce un valore `string` in cui la data in input è codificata nel formato **aaaa-mm-gg**.

##### Esempio d'esecuzione:

```text
$ cat test
2019/04/03
2019/6/4
2019/07/5
2019/4/05
$ go run date.go < test
2019-04-03
2019-06-04
2019-07-05
2019-04-05
```
