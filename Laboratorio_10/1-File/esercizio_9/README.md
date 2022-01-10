# Tragitti

Sul piano cartesiano, ad ogni punto individuato da una coppia di numeri reali, chiamati rispettivamente *ascissa* e *ordinata*, può essere associata un'*etichetta* simbolica, generalmente una lettera maiuscola.

Scrivere un programma che legga da **riga di comando** due stringhe di caratteri `Punti.txt` e `Tragitti.txt`. Ciascuna delle stringhe è il nome di un file di testo memorizzato nella stessa directory in cui è memorizzato il programma.

a) Ogni riga contenuta nel file `Punti.txt` è una stringa nel seguente formato:

*etichetta*;*x*;*y*
   
La tripla di valori separati dal carattere `;` specifica un punto sul piano cartesiano:
   
1. *etichetta*: una stringa che specifica l'etichetta simbolica associata al punto (ad es.: "A", "B", ...)
2. *x*: un valore reale che specifica l'ascissa del punto;
3. *y*: un valore reale che specifica l'ordinata del punto.

Esempio:

```text
A;10.0;2.0
B;11.5;3.0
C;8.0;1.0
```

b) Ogni riga contenuta nel file `Tragitti.txt` è una stringa nel seguente formato:

*id_tragitto*;*etichetta_punto_di_partenza*;*etichetta_punto_di_arrivo*
   
La tripla di valori separati dal carattere `;` specifica una tratta dal punto `etichetta_punto_di_partenza` al punto `etichetta_punto_di_arrivo` del tragitto `id_tragitto`.

Per ciascuna tripla, i valori `etichetta_punto_di_partenza` e `etichetta_punto_di_arrivo` sono diversi; in particolare, ciascun valore è associato ad una tripla memorizzata nel file `Punti.txt`.
La lunghezza di ciascuna tratta del tragitto, pari alla distanza euclidea tra i due punti che definiscono la tratta, è quindi ricavabile a partire dai dati memorizzati nel file `Punti.txt`.
Per esempio, la lunghezza della tratta:

*id_tragitto*;*A*;*B*

è pari alla distanza euclidea tra i punti `A` e `B`:
 
 ((x<sub>A</sub>-x<sub>B</sub>)<sup><small>2</small></sup> + (y<sub>A</sub>-y<sub>B</sub>)<sup><small>2</small></sup>)<sup><small>1/2</small></sup> = ((10.0-11.5)<sup><small>2</small></sup> + (2-3)<sup><small>2</small></sup>)<sup><small>1/2</small></sup>.

Le triple associate allo stesso tragitto `id_tragitto` specificano le tratte che definiscono il tragitto `id_tragitto`.

Due triple di valori:

*id_tragitto*;*etichetta_punto_di_partenza_A*;*etichetta_punto_di_arrivo_A*
*id_tragitto*;*etichetta_punto_di_partenza_B*;*etichetta_punto_di_arrivo_B*

specificano due tratte consecutive del tragitto `id_tragitto` se e solo se `etichetta_punto_di_arrivo_A` è uguale a `etichetta_punto_di_partenza_B`.

Ogni punto può comparire al più una volta nella sequenza di punti che definisce un dato tragitto.

Nel file, le triple di valori che specificano le diverse tratte di ciascun tragitto sono memorizzate in ordine casuale.

Il programma deve stampare a video la descrizione del tragitto più lungo definito dalle triple di valori memorizzate in `Tragitti.txt` (cfr. **Esempio di esecuzione**).

Si assuma che:
* le righe di testo contenute nei file `Punti.txt` e `Tragitti.txt` siano nel formato corretto;
* la tripla di valori presente in ogni riga del file `Punti.txt` specifichi correttamente un punto sul piano cartesiano;
* la tripla di valori presente in ogni riga del file `Tragitti.txt` specifichi correttamente la tratta di un tragitto;
* le triple di valori presenti nel file `Tragitti.txt` specifichino correttamente un insieme di tragitti non vuoto.

Potrebbero essere considerati i seguenti tipi di dato:

```go
type Punto struct {
	etichetta string
	x, y      float64
}

type Tratta struct {
	p1, p2 Punto
}

type TragittoPerTratte struct {
	id        int           // l'identificativo del tragitto
	tratte    []Tratta      // l'insieme delle tratte che definiscono il tragitto
	lunghezza float64       // la lunghezza del tragitto
}

type TragittoPerPunti struct {
	id        int           // l'identificativo del tragitto
	punti     []Punto       // la sequenza di punti che definisce il tragitto
	lunghezza float64       // la lunghezza del tragitto
}

```
Oltre alla funzione `main()`, tra le funzioni definite ed utilizzate potrebbero essere quindi incluse le seguenti funzioni:

* una funzione `Distanza(p1, p2 Punto) float64` che riceve in input due instanze del tipo `Punto` nei parametri `p1` e `p2` e restituisce un valore `float64` pari alla distanza euclidea tra i punti rappresentati da `p1` e `p2`;

* una funzione `NuovoTragittoPerTratte(identificativo int) (t *TragittoPerTratte)` che riceve in input un valore `int` nel parametro `identificativo` e restituisce un'istanza del tipo `TragittoPerTratte` nella variabile `t` i cui campi `id`, `tratte` e `lunghezza` sono  rispettivamente inizializzati ai valori `id`, `nil`  e `0.0`;

* una funzione `AggiungiTratta(t *TragittoPerTratte, p1, p2 Punto)` che riceve in input un'istanza del tipo `TragittoPerTratta` nel parametro `t` e due istanze del tipo `Punto` nei parametri `p1` e `p2`, e aggiunge l'instanza `Tratta{p1, p2}` del tipo `Tratta` all'insieme delle tratte che definiscono il tragitto rappresentato da `t`; 

* una funzione `NuovoTragittoPerPunti(TpT *TragittoPerTratte) (TpP *TragittoPerPunti)` che riceve in input un'istanza del tipo del tipo `TragittoPerTratte` nel papametro `TpT` e restituisce la corrispondente istanza del tipo `TragittoPerPunti` nella variabile `TpP` (i.e., `TpP.id = TpT.id`; `TpP.lunghezza = TpT.lunghezza`; le istanze del tipo `Punto` `TpP.punti[0]` e `TpP.punti[1]` definiscono la prima tratta del tragitto, le istanze `TpP.punti[1]` e `TpP.punti[2]` definiscono la seconda tratta del tragitto,...)
 
##### Esempio d'esecuzione:

```text
$ cat punti.txt 
A;0;0
B;3;5
C;1;4
D;-1;0
E;4;6
F;-4;-2
G;4;-5

$ cat tragitti.txt 
1;A;B
3;F;G
2;C;A
1;B;F
2;B;C
1;F;D
2;A;C

$ go run tragitti.go punti.txt tragitti.txt
Tragitto 1 - Lunghezza 19.335998106920954 A(0, 0) B(3, 5) F(-4, -2) D(-1, 0)
Tragitto 3 - Lunghezza 8.54400374531753 F(-4, -2) G(4, -5)
Tragitto 2 - Lunghezza 7.35917360311745 B(3, 5) C(1, 4) A(0, 0) D(-1, 0)

Tragitto di lunghezza massima: 
Tragitto 1 - Lunghezza 19.335998106920954 A(0, 0) B(3, 5) F(-4, -2) D(-1, 0)
```