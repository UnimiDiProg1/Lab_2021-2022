# Quadrato a colonne alterne (2)

Scrivere un programma che legga da **standard input** un numero intero `n` e che, come mostrato nell'**Esempio di esecuzione**, stampi a video un quadrato di `n` righe costituite ciascuna da `n`
simboli intervallati da spazi, alternando fra loro 2 colonne costituite da simboli `*` (asterisco) a 2 colonne
costituite da simboli `+` (più). In particolare, se `n` è dispari solo una delle due colonne più a destra sarà stampata.

##### Esempio d'esecuzione:

```text
$ go run quadrato_colonne_alterne_2.go
Inserisci un numero: 7
* * + + * * +
* * + + * * +
* * + + * * +
* * + + * * +
* * + + * * +
* * + + * * +
* * + + * * +

$ go run quadrato_colonne_alterne_2.go
Inserisci un numero: 5
* * + + *
* * + + *
* * + + *
* * + + *
* * + + *
```