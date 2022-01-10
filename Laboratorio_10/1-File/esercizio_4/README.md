# Filtra testo

Scrivere un programma che legga da **riga di comando** due stringhe di caratteri `file.txt` e `pattern`.

`file.txt` è il nome di un file di testo memorizzato nella stessa directory in cui è memorizzato il programma.

Il programma deve stampare a video le righe del file `file.txt` in cui occorre la stringa di caratteri `pattern`.

Il controllo della presenza del `pattern` in un riga di testo deve effettuato in modalità **case sensitive**: 
* Se il `pattern` fosse `sta`, il `pattern` occorre nella riga di testo `Dove sta andando?`;
* Se il `pattern` fosse `sta`, il `pattern` occorre nella riga di testo `Cosa stai facendo?`; 
* Se il `pattern` fosse `STA`, il `pattern` **non** occorre nella riga di testo `Dove sta andando?`; 

Si assuma che la stringa `pattern` specificata a riga di comando includa almeno un carattere.

*Suggerimento:* Per controllare la presenza di un pattern all'interno di una stringa potete utilizzare la funzione `strings.Contains` dal package `strings`.

##### Esempio d'esecuzione:

```text
$ go run filtra.go testo.txt sotto
e sotto il maestrale

$ go run filtra.go testo.txt tra  
e sotto il maestrale
tra le rossastre nubi

$ go run filtra.go testo.txt fischi
sta il cacciator fischiando

$ go run filtra.go testo.txt FISCHI

$ go run filtra.go testo.txt nube  
``` 
