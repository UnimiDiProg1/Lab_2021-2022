# Mix di due sequenze ordinate di numeri

Scrivere un programma che legga da **riga di comando** due stringhe di caratteri `sequenza_ordinata_1.txt` e `sequenza_ordinata_2.txt`. Ciascuna delle stringhe è il nome di un file di testo memorizzato nella stessa directory in cui è memorizzato il programma, il cui contenuto è una sequenza di numeri interi separati dal carattere spazio (` `) e ordinati dal più piccolo al più grande.

Siano `s_1` e `s_2` le sequenze ordinate di numeri memorizzate rispettivamente nei file `sequenza_ordinata_1.txt` e `sequenza_ordinata_1.txt`.

Il programma deve stampare su un nuovo file di testo `sequenza_risultante.txt` la sequenza ordinata di numeri `s` ottenibile dal mix delle sequenze `s_1` e `s_2`: 
* `s` è una sequenza che include tutti i numeri presenti in `s_1` e `s_2`;
* `s` è una sequenza di numeri interi separati dal carattere spazio (` `) e ordinati dal più piccolo al più grande.

##### Esempio d'esecuzione:

```text
$ cat sequenza_ordinata_1.txt 
1 5 8 10 34 56

$ cat sequenza_ordinata_2.txt
2 4 7 9 10 23 45

$ go run sequenze_ordinate.go sequenza_ordinata_1.txt sequenza_ordinata_2.txt
Execution time: 0.0009968

$ cat sequenza_risultante.txt                                                  
1 2 4 5 7 8 9 10 10 23 34 45 56

$ go run sequenze_ordinate.go sequenza_ordinata_A.txt sequenza_ordinata_B.txt
Execution time: 7.8574857
``` 