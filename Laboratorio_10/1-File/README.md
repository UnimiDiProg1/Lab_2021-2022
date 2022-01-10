# Laboratorio 10 - File
## 1 Qual è l'output?

Supponendo che:
* l'utente inserisca da **riga di comando** il valore `punti.txt`,
* il file `punti.txt` contenga il seguente testo:
```text
A 10.5 20
B 15 30
C 12.5 25.6
```
* il file `punti.txt` sia memorizzato nella stessa directory in cui è memorizzato il programma di seguito riportato,

cosa dovrebbe produrre in output il programma? 

```go
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("Error while opening the file! %v\n", err)
		return
	}
	defer f.Close()

	for {
		var nome string
		var x, y float64
		_, err = fmt.Fscan(f, &nome, &x, &y)
        // ... equivalente a:
        // _, err = fmt.Fscanf(f, "%s%f%f", &nome, &x, &y)
        // ... oppure a:
        // _, err = fmt.Fscanf(f, "%s %f %f", &nome, &x, &y)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Error while reading the file! %v\n", err)
			return
		}
		fmt.Printf("Punto %s = (%v, %v)\n", nome, x, y)
	}
}
```

## 2 Qual è l'output?

Supponendo che l'utente inserisca da **riga di comando** il valore `tabellina.txt`, cosa dovrebbe produrre in output il seguente programma?

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create(os.Args[1])
	if err != nil {
		fmt.Printf("Error while creating the file! %v\n", err)
		return
	}
	defer f.Close()

	for i:=1; i<=10; i++ {
		risultato := i * 10
		_, err = fmt.Fprintln(f, "10 x ", i, risultato)
        // ... equivalente a:
        // _, err = fmt.Fprintf(f, "10 x %d = %d\n", i, risultato)
		if err != nil {
			fmt.Printf("Error while writing the file! %v\n", err)
			return
		}
	}
}
```

## 3 Qual è l'output?

Supponendo che:
* l'utente inserisca da **riga di comando** i valori `testo.in testo.out` ,
* il file `testo.in` contenga il seguente testo:
```text
La nebbia a gl'irti colli
piovigginando sale,
e sotto il maestrale
urla e biancheggia il mar;
```
* il file `testo.in` sia memorizzato nella stessa directory in cui è memorizzato il programma di seguito riportato,

cosa dovrebbe produrre in output il programma? 

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		fIn, fOut *os.File
		err error
	)

	fIn, err = os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("Error while opening the file! %v\n", err)
		return
	}
	defer fIn.Close()
	fOut, err = os.Create(os.Args[2])
	if err != nil {
		fmt.Printf("Error while creating the file! %v\n", err)
		return
	}
	defer fOut.Close()

	scanner := bufio.NewScanner(fIn)
	counter := 0
	for scanner.Scan() {
		line := scanner.Text()
		counter++
		_, err = fmt.Fprintf(fOut,"LINE %3d: <%s>\n", counter, line)
		if err != nil {
			fmt.Printf("Error while writing the file! %v\n", err)
			return
		}
	}
	if err = scanner.Err(); err != nil {
		fmt.Printf("Error while reading the file! %v\n", err)
		return
	}
}
```

## 4 Filtra testo

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

## 5 Potenze

Scrivere un programma che legga da **riga di comando** due numeri interi `b` ed `e` e stampi su un nuovo file di testo `potenze.txt` le potenze del numero `b` fino a `b`<sup>`e`</sup>.

*Esempio*: Se `b=2` ed `e=5`, il programma deve stampare sul file `potenze.txt` i numeri `1 2 4 8 16 32`.

##### Esempio d'esecuzione:

```text
$ go run potenze.go 2 5

$ cat potenze.txt      
1 2 4 8 16 32

$ go run potenze.go 4 3

$ cat potenze.txt      
1 4 16 64
```
## 6 Ripetizioni


Scrivere un programma che legga da **riga di comando** una stringa di caratteri `file.txt`.

`file.txt` è il nome di un file di testo memorizzato nella stessa directory in cui è memorizzato il programma.

Il programma deve stampare su un nuovo file di testo `statistiche.txt`: 
1. Il numero di parole distinte presenti nel file `file.txt` (una parola è una stringa interamente definita da caratteri il cui codice Unicode, se passato come argomento alla funzione `func IsSpace(r rune) bool` del package `unicode`, fa restituire `false` alla funzione).
2. La lista di parole distinte presenti nel file `file.txt`, riportando per ogni parola ripetuta il relativo numero di occorrenze (cfr. **Esecuzione d'esecuzione**).

##### Esempio d'esecuzione:

```text
$ cat testo.txt 
Lontano, nei dimenticati spazi non segnati nelle carte geografiche 
dell'estremo limite della Spirale Ovest della Galassia, c'è un piccolo 
e insignificante sole giallo.
A orbitare intorno a esso, alla distanza di centoquarantanove milioni 
di chilometri, c'è un piccolo, trascurabilissimo pianeta azzurro–verde, 
le cui forme di vita, discendenti dalle scimmie, sono così incredibilmente
primitive che credono ancora che gli orologi da polso digitali siano
un'ottima invenzione.

Guida galattica per gli autostoppisti

$ go run ripetizioni.go testo.txt

$ more statistiche.txt
Totale parole distinte: 62
Occorrenze:
LONTANO,: 1
DELLA: 2
DALLE: 1
CHE: 2
SEGNATI: 1
DELL'ESTREMO: 1
A: 2
CHILOMETRI,: 1
PIANETA: 1
UN'OTTIMA: 1
COSÌ: 1
--More--(21%)
``` 
## 7 Mix di due sequenze ordinate di numeri

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
## 8 Mix di n sequenze ordinate di numeri

Scrivere un programma che legga da **riga di comando** una stringa di caratteri `lista_file.txt`.

`lista_file.txt` è il nome di un file di testo memorizzato nella stessa directory in cui è memorizzato il programma, il cui contenuto è una sequenza di stringhe separate dal carattere spazio (`;`):
 `sequenza_ordinata_1.txt;sequenza_ordinata_2.txt;...;sequenza_ordinata_n.txt`.
Ciascuna delle stringhe presenti in `lista_file.txt` è a sua volta il nome di un file di testo memorizzato nella stessa directory in cui è memorizzato il programma, il cui contenuto è una sequenza di numeri interi separati dal carattere spazio (` `) e ordinati dal più piccolo al più grande.

Il programma deve stampare su un nuovo file di testo `sequenza_risultante.txt` la sequenza ordinata di numeri ottenibile dal mix delle sequenze memorizzate nei file il cui nome è memorizzato in `lista_file.txt`.

##### Esempio d'esecuzione:

```text
$ cat lista_file.txt 
sequenza_ordinata_1.txt;sequenza_ordinata_2.txt;sequenza_ordinata_3.txt

$ cat sequenza_ordinata_1.txt 
1 5 8 10 34 56
                   
$ cat sequenza_ordinata_2.txt
2 4 7 9 10 23 45
                   
$ cat sequenza_ordinata_3.txt
4 7 9 10 22

$ go run sequenze_ordinate.go lista_file_i.txt
Execution time: 0.0019944

$ cat sequenza_risultante.txt               
1 2 4 4 5 7 7 8 9 9 10 10 10 22 23 34 45 56 

$ go run sequenze_ordinate.go lista_file_ii.txt
Execution time: 37.5841148
``` 
## 9 Tragitti

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
## 10 Lettura di file 'PPM'

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
## 11 Scrittura di file 'PPM'

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
## 12 Copia di file 'PPM'

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
## 13 Conversione di file 'PPM'

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
