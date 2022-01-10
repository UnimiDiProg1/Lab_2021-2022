# Mix di n sequenze ordinate di numeri

Scrivere un programma che legga da **riga di comando** una stringa di caratteri `lista_file.txt`.

`lista_file.txt` è il nome di un file di testo memorizzato nella stessa directory in cui è memorizzato il programma, il cui contenuto è una sequenza di stringhe separate dal carattere spazio (`;`):
 `sequenza_ordinata_1.txt;sequenza_ordinata_2.txt;...;sequenza_ordinata_n.txt`.
Ciascuna delle stringhe presenti in `lista_file.txt` è a sua volta il nome di un file di testo memorizzato nella stessa directory in cui è memorizzato il programma, il cui contenuto è una sequenza di numeri interi separati dal carattere spazio (` `) e ordinati dal più piccolo al più grande.

Il programma deve stampare su un nuovo file di testo `sequenza_risultante.txt` la sequenza ordinata di numeri ottenibile dal mix delle sequenze memorizzate nei file il cui nome è memorizzato in `lista_file.txt`.

##### Esempio d'esecuzione:

```text
$ cat lista_file.txt 
sequenza_ordinata_1.txt;sequenza_ordinata_2.txt;sequenza_ordinata_3.txt

$ cat sequenza_ordinata_1.txt 
1 5 8 10 34 56
                   
$ cat sequenza_ordinata_2.txt
2 4 7 9 10 23 45
                   
$ cat sequenza_ordinata_3.txt
4 7 9 10 22

$ go run sequenze_ordinate.go lista_file_i.txt
Execution time: 0.0019944

$ cat sequenza_risultante.txt               
1 2 4 4 5 7 7 8 9 9 10 10 10 22 23 34 45 56 

$ go run sequenze_ordinate.go lista_file_ii.txt
Execution time: 37.5841148
``` 