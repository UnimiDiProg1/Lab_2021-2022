# Classifica triangolo

Scrivere un programma che legga da **standard input** sei valori reali:
- il primo ed il secondo valore, `xA` e `yA`, rappresentano rispettivamente l'ascissa e l'ordinata di un punto `A` sul piano cartesiano;
- il terzo ed il quarto valore, `xB` e `yB`, rappresentano rispettivamente l'ascissa e l'ordinata di un punto `B` sul piano cartesiano;
- il quinto ed il sesto valore, `xC` e `yC`, rappresentano rispettivamente l'ascissa e l'ordinata di un punto `C` sul piano cartesiano.

Una volta terminata la fase di lettura, il programma deve stampare a video (come mostrato nell'**Esempio di esecuzione**), se il triangolo `ABC` definito dai segmenti/lati `AB`, `BC` e `AC` è equilatero, iscoscele o scaleno.

Un triangolo è equilatero se ha tutti e tre i lati di lunghezza uguale.

Un triangolo è isoscele se ha solo due lati di lunghezza uguale.

Un triangolo è scaleno se ha tutti e tre i lati di lunghezza diversa.

La lunghezza di ciascun lato di un triangolo è pari alla distanza euclidea tra gli estremi del lato.
 
Per esempio, la lunghezza del lato `AB` del triangolo `ABC` è pari alla distanza euclidea tra i punti `A` e `B`: ((xA-xB)<sup><small>2</small></sup> + (yA-yB)<sup><small>2</small></sup>)<sup><small>1/2</small></sup>.

Le lunghezze dei lati del triangolo vanno confrontate a meno di una costante `EPSILON = 1.0e-9`.

Oltre alla funzione `main()`, devono essere definite ed utilizzate almeno le seguenti funzioni:

* una funzione `Distanza(x1, y1, x2, y2 float64) float64` che riceve in input:

    - due valori `float64` nei parametri `x1` e `y1` che rappresentano rispettivamente l'ascissa e l'ordinata di un punto `P1` sul piano cartesiano;
    
    - due valori `float64` nei parametri `x2` e `y2` che rappresentano rispettivamente l'ascissa e l'ordinata di un punto `P2` sul piano cartesiano;

    e restituisce un valore `float64` pari alla distanza euclidea tra i punti `P1` e `P2`.

##### Esempio d'esecuzione:

```text
$ go run classifica.go
Inserisci i valori dell'ascissa e dell'ordinata del punto A: -6 1
Inserisci i valori dell'ascissa e dell'ordinata del punto B: -1 1
Inserisci i valori dell'ascissa e dell'ordinata del punto C: -3.5 5.330127018922193

Il triangolo ABC è equilatero.
Lunghezza del lato: 5

$ go run classifica.go
Inserisci i valori dell'ascissa e dell'ordinata del punto A: -6 1
Inserisci i valori dell'ascissa e dell'ordinata del punto B: -1 1
Inserisci i valori dell'ascissa e dell'ordinata del punto C: -3.5 5.33

Il triangolo ABC è isocele.
I lati di lunghezza uguale sono AC e BC.
Lunghezza dei lati AC e BC: 4.999889998789973
Lunghezza del lato AB: 5

$ go run classifica.go
Inserisci i valori dell'ascissa e dell'ordinata del punto A: 1 1
Inserisci i valori dell'ascissa e dell'ordinata del punto B: 3 3
Inserisci i valori dell'ascissa e dell'ordinata del punto C: 1 5

Il triangolo ABC è isocele.
I lati di lunghezza uguale sono AB e BC.
Lunghezza dei lati AB e BC: 2.8284271247461903
Lunghezza del lato AC: 4

$ go run classifica.go
Inserisci i valori dell'ascissa e dell'ordinata del punto A: 1 1
Inserisci i valori dell'ascissa e dell'ordinata del punto B: 1 5
Inserisci i valori dell'ascissa e dell'ordinata del punto C: 5 1

Il triangolo ABC è isocele.
I lati di lunghezza uguale sono AB e AC.
Lunghezza dei lati AB e AC: 4
Lunghezza del lato BC: 5.656854249492381

$ go run classifica.go
Inserisci i valori dell'ascissa e dell'ordinata del punto A: 1 1
Inserisci i valori dell'ascissa e dell'ordinata del punto B: 1 5
Inserisci i valori dell'ascissa e dell'ordinata del punto C: 3 3

Il triangolo ABC è isocele.
I lati di lunghezza uguale sono AC e BC.
Lunghezza dei lati AC e BC: 2.8284271247461903
Lunghezza del lato AB: 4

$ go run classifica.go
Inserisci i valori dell'ascissa e dell'ordinata del punto A: 1 1
Inserisci i valori dell'ascissa e dell'ordinata del punto B: 1 4
Inserisci i valori dell'ascissa e dell'ordinata del punto C: 1.5 1.5

Il triangolo ABC è scaleno.
Lunghezza del lato AB: 3
Lunghezza del lato BC: 2.5495097567963922
Lunghezza del lato AC: 0.7071067811865476
``` 

