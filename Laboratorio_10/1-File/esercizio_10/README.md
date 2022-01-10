# Lettura di file 'PPM'

<sub><sup>*I concetti di seguito illustrati sono solo un'esemplificazione dei corrispondenti concetti reali. 
Le definizioni fornite hanno il solo scopo di permettere lo svolgimento degli esercizi; in particolare, la fomulazione di definizioni precise e rigorose per i concetti introdotti va al di là degli scopi di queste dispense.*</sup></sub>

I file `PPM` sono file di testo in cui sono memorizzati (in formato testuale) i dati che definiscono un'immagine scomposta in pixel.
Il pixel è il più piccolo elemento grafico visibile sullo schermo, corrispondente a un punto luminoso; ciascuno dei punti di cui si compone un'immagine digitale.

`test.ppm` è un esempio di file in formato `PPM`:
```text
$ cat test.ppm
P3
4 2
240
200 100 100
100 200 200
240 240 240
130 130 145
100 100 100
80 100 120
40 60 50
10 20 30
```
* `P3` (prima riga) indica che le triple di interi specificate in modo testuale su ciascuna riga del file, dalla quarta in poi, indicano il valore di rosso, verde e blue da considerare per la colorazione di un pixel dell'immagine in base al modello RGB. RGB è un modello di colori di tipo additivo che permette di specificare il colore di un pixel come somma dei tre colori Rosso (Red), Verde (Green) e Blu (Blue).
* 4 e 2 (seconda riga) indicano che l'immagine è larga 4 pixel e alta 2 pixel:
    1. le righe dalla quarta alla settima specificano la colorazione dei 4 pixel che definiscono la parte più alta dell'immagine, da sinistra verso destra guardando l'immagine;
    1. le righe dall'ottava all'undicesima specificano la colorazione dei 4 pixel che definiscono la parte più bassa dell'immagine, da sinistra verso destra guardando l'immagine.
* 240 (terza riga) è il massimo valore specificato per un colore di un pixel dell'immagine.


Scrivere un programma che legga da **riga di comando** una stringa di caratteri `immagine.ppm`
 
 `immagine.ppm` è il nome di un file `PPM` memorizzato nella stessa directory in cui è memorizzato il programma.

In base ai dati memorizzati nel file `immagine.ppm`, il programma deve inizializzare in modo opportuno un'istanza del tipo `Immagine` definito come di seguito:

```go
type RGB struct {
    r, g, b uint
}
type Immagine struct {
    larghezza, altezza uint
    valori []RGB
}
``` 
In particolare, il programma deve:
1. Aprire il file `immagine.ppm`.
2. Dichiarare una nuova variabile `immagine` di tipo `Immagine`.
2. Leggere dal file un valore di tipo `string` (`P3` nell'esempio relativo al file `test.ppm`) e ignorarlo.
3. Leggere dal file due valori `uint` che rappresentano la larghezza e l'altezza dell'immagine e memorizzarli nei relativi campi di `immagine`.
4. Leggere dal file un valore `uint` e ignorarlo.
5. Leggere iterativamente 3 valori `uint` dal file e, ad ogni iterazione:
    1. Salvare i tre valori nei relativi campi di una variabile di tipo `RGB`.
    2. Aggiungere il valore della variabile di tipo `RGB` al campo `valori` di `immagine`.
6. Stampare il contenuto di `immagine`.
 
Oltre alla funzione `main()`, devono essere definite ed utilizzate almeno le seguenti funzioni:
* una funzione `LeggiFilePPM(fileInput *os.File) (immagine Immagine, err error)` che riceve in input, nel parametro `fileInput`, un'istanza del tipo `os.File` relativa ad un file `PPM` in cui sono memorizzati (in formato testuale) i dati che definiscono un'immagine scomposta in pixel, e restituisce:
    1. un'istanza del tipo `Immagine` nella variabile `immagine` inizializzata in base ai dati memorizzati nel file `PPM`;  
    2. un valore `error` nella variabile `err` pari a `nil` se la lettura del file è andata a buon fine (non ha generato errori), o al valore generato da una delle funzione utilizzate per la lettura del file altrimenti.      
     
##### Esempio d'esecuzione:

```text
$ go run image.go test.ppm 
Immagine letta: {4 2 [{200 100 100} {100 200 200} {240 240 240} {130 130 145} {100 100 100} {80 100 120} {40 60 50} {10 20 30}]}
```