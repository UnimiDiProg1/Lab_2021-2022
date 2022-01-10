# Conversione di file 'PPM'

Scrivere un programma che legga da **riga di comando** due stringhe di caratteri `immagine.ppm` e `immagine_scala_di_grigi.ppm`.
 
`immagine.ppm` è il nome di un file `PPM` in cui sono memorizzati (in formato testuale) i dati che definiscono un'immagine a *colori* scomposta in pixel, memorizzato nella stessa directory in cui è memorizzato il programma.

`immagine_scala_di_grigi.ppm` è il nome di un nuovo file `PPM` da creare nella stessa directory in cui è memorizzato il programma.

Il programma deve leggere il contenuto del file `immagine.ppm` e stampare sul nuovo file `immagine_scala_di_grigi.ppm` i dati che definiscono un'immagine a *scala di grigi* scomposta in pixel **uguale** all'immagine a *colori* definita dai dati memorizzati nel file `immagine.ppm`, **eccezion fatta** per la colorazione: i valori `rosso`, `verde` e `blue` da considerare per la colorazione di un pixel dell'immagine a *colori* devono essere sostituiti rispettivamente dai valori `rosso'`, `verde'` e `blue'` per la colorazione del corrispondente pixel dell'immagine a *scala di grigi*, con: 

`rosso' = verde' = blue' = (rosso + verde + blue)/3`

Oltre alla funzione `main()`, devono essere definite ed utilizzate almeno le seguenti funzioni:

* una funzione `LeggiFilePPM(fileInput *os.File) (immagine Immagine, err error)` che riceve in input, nel parametro `fileInput`, un'istanza del tipo `os.File` relativa ad un file `PPM` in cui sono memorizzati (in formato testuale) i dati che definiscono un'immagine scomposta in pixel, e restituisce:
    1. un'istanza del tipo `Immagine` nella variabile `immagine` inizializzata in base ai dati memorizzati nel file `PPM`;  
    2. un valore `error` nella variabile `err` pari a `nil` se la lettura del file è andata a buon fine (non ha generato errori), o al valore generato da una delle funzione utilizzate per la lettura del file altrimenti;

* una funzione `ConvertiImmagine(immagineAColori Immagine) (immagineAScalaDiGrigi Immagine)` che riceve in input un'istanza del tipo `Immagine` nel parametro `immagineAColori` (in cui sono memorizzati i dati che definiscono un'immagine a *colori* scomposta in pixel) e restituisce un'istanza del tipo `Immagine` nel parametro `immagineAScalaDiGrigi` in cui sono memorizzati i dati che definiscono un'immagine a *scala di grigi* scomposta in pixel **uguale** all'immagine a *colori* definita dai dati memorizzati in `immagineAColori`, **eccezion fatta** per la colorazione: nei campi `r`, `g`  e `b` dell'i-esimo valore di `immagineAScalaDiGrigi.valori` deve essere memorizzato il valore della media aritmetica, arrotondata all'intero, dei valori dei campi `r`, `g`  e `b` dell'i-esimo valore di `immagineAColori.valori`; 

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

$ go run image.go test.ppm test_bn.ppm 

$ cat test_bn.ppm 
P3
4 2
240
133 133 133
166 166 166
240 240 240
135 135 135
100 100 100
100 100 100
50 50 50
20 20 20
```