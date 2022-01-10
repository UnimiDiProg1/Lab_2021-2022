# Copia di file 'PPM'

Scrivere un programma che legga da **riga di comando** due stringa di caratteri `immagine.ppm` e `copia.ppm`.
 
`immagine.ppm` è il nome di un file `PPM` memorizzato nella stessa directory in cui è memorizzato il programma.

`copia.ppm` è il nome di un nuovo file `PPM` da creare nella stessa directory in cui è memorizzato il programma.

Il programma deve leggere il contenuto del file `immagine.ppm` e stamparlo sul nuovo file `copia.ppm`.

Oltre alla funzione `main()`, devono essere definite ed utilizzate almeno le seguenti funzioni:
* una funzione `LeggiFilePPM(fileInput *os.File) (immagine Immagine, err error)` che riceve in input, nel parametro `fileInput`, un'istanza del tipo `os.File` relativa ad un file `PPM` in cui sono memorizzati (in formato testuale) i dati che definiscono un'immagine scomposta in pixel, e restituisce:
    1. un'istanza del tipo `Immagine` nella variabile `immagine` inizializzata in base ai dati memorizzati nel file `PPM`;  
    2. un valore `error` nella variabile `err` pari a `nil` se la lettura del file è andata a buon fine (non ha generato errori), o al valore generato da una delle funzione utilizzate per la lettura del file altrimenti;
* una funzione `ScriviFilePPM(fileOutput *os.File, immagine Immagine) (err error)` che riceve in input:
    1. un'istanza del tipo `Immagine` nel parametro `immagine` (in cui sono memorizzati i dati che definiscono un'immagine scomposta in pixel);
    2. un'istanza del tipo `os.File` nel parametro `fileOutput` relativa ad un file `PPM` in cui memorizzare (in formato testuale) i dati che definiscono un'immagine scomposta in pixel; 
    
    e restituisce  un valore `error` nella variabile `err` pari a `nil` se la scrittura del file è andata a buon fine (non ha generato errori), o al valore generato da una delle funzione utilizzate per la scrittura del file altrimenti.

##### Esempio d'esecuzione:

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

$ go run image.go test.ppm copia.ppm 

$ cat copia.ppm 
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