# Albero

Scrivere un programma che legga da **standard input** un intero `n` e, come mostrato nell'**Esempio d'esecuzione**, stampi a video un albero utilizzando il carattere `*` (asterisco) per rappresentarne la chioma ed il carattere `+` (più) per rappresentarne il tronco:
* La chioma dell'albero deve essere alta `n` righe e, nel punto di larghezza massima, larga `2*n-1` colonne.
* Il tronco dell'albero deve essere rappresentato con un quadrato di altezza e larghezza pari a `3` caratteri. 
* Il tronco dell'albero deve essere centrato rispetto alla chioma dell'albero.

Se `n<=0`, il programma stampa solo il tronco dell'albero.

##### Esempio d'esecuzione:

```text
$ go run albero.go
7
      *
     ***
    *****
   *******
  *********
 ***********
*************
     +++
     +++
     +++

$ go run albero.go
4
   *
  ***
 *****
*******
  +++
  +++
  +++

$ go run albero.go
1
 *
+++
+++
+++

$ go run albero.go
0
+++
+++
+++

$ go run albero.go
-1
+++
+++
+++
```