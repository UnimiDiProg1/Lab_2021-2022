# Scrittura di file 'PPM'

Scrivere un programma che legga da **riga di comando** una stringa di caratteri `immagine.ppm`.
  
Il programma deve stampare su un nuovo file di testo `immagine.ppm` i dati memorizzati in un'istanza del tipo `Immagine` (vedi Esercizio 10 (Laboratorio 12 - File)). I dati devono essere opportunamente stampati affinché `immagine.ppm` sia un file in formato `PPM` (vedi Esercizio 10 (Laboratorio 12 - File)). 

In particolare, il programma deve:
1. Creare il file `immagine.ppm`.
2. Dichiarare una nuova variabile `immagine` di tipo `Immagine`.
3. Inizializzare la variabile `immagine` come segue:
   ```go
   immagine.larghezza = 4
   immagine.altezza = 2
   immagine.valori = []RGB{RGB{200, 100, 100}, RGB{100, 200, 200},
   		RGB{240, 240, 240}, RGB{130, 130, 145}, RGB{100, 100, 100},
   		RGB{80, 100, 120}, RGB{40, 60, 50}, RGB{10, 20, 30}}
   ```
3. Scrivere sulla prima riga del file la stringa `"P3"`.
3. Scrivere sulla seconda riga del file i valori dei campi `larghezza` e `altezza` di `immagine` (separati dal carattere ` `)
4. Scrivere sulla terza riga del file il massimo tra tutti i valori `uint` di tutti i valori `RGB` memorizzati nel campo `valori` di `immagine`;
5. Scrivere sul file una riga aggiuntiva per ogni valore `RGB` memorizzato nel campo `valori` di `immagine` (dal primo valore `RGB` all'ultimo valore `RGB`); su ogni riga devono comparire i valori dei campi `r`, `g` e `b`  (separati dal carattere ` `) del corrispondete valore `RGB`.

Oltre alla funzione `main()`, devono essere definite ed utilizzate almeno le seguenti funzioni:
* una funzione `ScriviFilePPM(fileOutput *os.File, immagine Immagine) (err error)` che riceve in input:
    1. un'istanza del tipo `Immagine` nel parametro `immagine` (in cui sono memorizzati i dati che definiscono un'immagine scomposta in pixel);
    2. un'istanza del tipo `os.File` nel parametro `fileOutput` relativa ad un file `PPM` in cui memorizzare (in formato testuale) i dati che definiscono un'immagine scomposta in pixel; 
    
    e restituisce  un valore `error` nella variabile `err` pari a `nil` se la scrittura del file è andata a buon fine (non ha generato errori), o al valore generato da una delle funzione utilizzate per la scrittura del file altrimenti.

* una funzione `MassimoValore(valori []RGB) uint` che riceve in input un valore `[]RGB` nel parametro `valori` e restituisce un valore `uint` pari al massimo tra tutti i valori `uint` di tutti i valori `RGB` memorizzati nel campo `valori`.

##### Esempio d'esecuzione:

```text
$ go run image.go test.ppm 

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