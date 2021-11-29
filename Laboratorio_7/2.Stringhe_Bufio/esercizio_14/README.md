# Date (2)

Scrivere un programma che:
* legga da **standard input** un testo su più righe;
* termini la lettura quando viene inserita da standard input una riga vuota ( "" ).

Ogni riga di testo è una una stringa senza spazi che codifica una data in uno dei seguenti possibili formati:

1. aaaa/m/g
2. aaaa/m/gg
3. aaaa/mm/g
4. aaaa/mm/gg
5. g/m/aaaa
6. gg/m/aaaa
7. g/mm/aaaa
8. gg/mm/aaaa

Una volta terminata la fase di lettura, il programma deve stampare la codifica nel formato aaaa-mm-gg delle date lette. In particolare, le date devono essere stampate in ordine cronologico (dalla più antica alla più recente).

Oltre alla funzione `main()` deve essere definita ed utilizzata la funzione `Formatta(data string) string`, che data una data in input codificata in uno degli 8 formati descritti sopra, la restituisce formattata aaaa-mm-gg. La funzione deve utilizzare `strings.Split` per estrarre giorno, mese, e anno dalla stringa di input.

*Suggerimento:* Una volta codificate nel formato **aaaa-mm-gg**, le date possono essere ordinate cronologicamente utilizzando la funzione `sort.Strings(a []string)` del package `sort`.

##### Esempio d'esecuzione:

```text
$ cat test
2019/04/03
2019/6/4
2019/07/5
2019/4/05
5/5/2019
05/4/2019
7/05/2019
07/09/2019
07/09/2018

$ go run date.go < test      
2018-09-07
2019-04-03
2019-04-05
2019-04-05
2019-05-05
2019-05-07
2019-06-04
2019-07-05
2019-09-07
```
