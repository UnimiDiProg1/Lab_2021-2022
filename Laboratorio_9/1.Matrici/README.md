# Qual è l'output?

Qual è l'output del seguente programma?

```go
package main

import (
	"fmt"
)

func main() {

	var a [3][2]int

	fmt.Printf("Tipo di a: %T: %v\n", a, a)

	a = [3][2]int{{1, 2}, {10, 20}, {100, 200}}

	fmt.Printf("Tipo di a: %T: %v\n", a, a)

	for i := 0; i < len(a); i++ {
		fmt.Printf("a[%d] - tipo: %T: %v\n", i, a[i], a[i])
		for j := 0; j < len(a[i]); j++ {
			fmt.Printf("a[%d][%d] - tipo: %T: %v\n", i, j, a[i][j], a[i][j])
		}
	}
}
```
# Qual è l'output?

Qual è l'output del seguente programma?

```go
package main

import (
	"fmt"
)

func main() {

	var a [][]int

	fmt.Printf("Tipo di a: %T: %v\n", a, a)

	a = [][]int{{1, 2}, {10, 20}, {100, 200}}

	fmt.Printf("Tipo di a: %T: %v\n", a, a)

	for i := 0; i < len(a); i++ {
		fmt.Printf("a[%d] - tipo: %T: %v\n", i, a[i], a[i])
		for j := 0; j < len(a[i]); j++ {
			fmt.Printf("a[%d][%d] - tipo: %T: %v\n", i, j, a[i][j], a[i][j])
		}
	}
}
```
# Qual è l'output?

Qual è l'output del seguente programma?

```go
package main

import (
	"fmt"
)

func main() {

	var n int = 5
	var m int = 4

	var a [][]int

	a = make([][]int,n)

	for i := 0; i < n; i++ {
		a_i := make([]int,m)
		
		for j := i; j < m+i; j++ {
			a_i = append(a_i,j)
			fmt.Printf("%d ",j)
		}
		
		a = append(a, a_i)
		fmt.Println()
	}
}
```
# Qual è l'output?

Qual è l'output del seguente programma?

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	n := 4
	var slc [][]int = CreaSliceBidimensionale(n)
	AssegnaSliceBidimensionale(slc)
	fmt.Println(slc)

	slc_soprasoglia := FiltraSliceBidimensionale(slc, 2)
	fmt.Println(slc_soprasoglia)
}

func CreaSliceBidimensionale(l int) [][]int {
	var s [][]int = make([][]int, l)
	for i := 0; i < l; i++ {
		s[i] = make([]int, l)
	}
	return s
}

func AssegnaSliceBidimensionale(s [][]int) {
	rand.Seed(int64(time.Now().Nanosecond()))
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s[i]); j++ {
			s[i][j] = rand.Intn(2)
		}
	}
}

func FiltraSliceBidimensionale(s [][]int, soglia int) [][]int {
	var risultato [][]int
	for _,riga := range s {
		somma := 0
		for _,valore := range riga {
			somma += valore
		}
		if somma >= soglia {
			risultato = append(risultato, riga)
		}
	}
	return risultato
}
```
# Tavola pitagorica

Scrivere un programma che legga da **riga di comando** un numero intero `n` e, come mostrato nell'**Esempio di esecuzione**, stampi a video la corrispondente tavola pitagorica `n x n`.

Oltre alla funzione `main()`, devono essere definite ed utilizzate almeno le seguenti funzioni:

* una funzione `CreaTavolaPitagorica(n int) [][]int` che riceve in input un valore `int` nel parametro `n` e restituisce un valore di tipo `[][]int` in cui sono memorizzati i valori di una tavola pitagorica `n x n`;
* una funzione `StampaTavolaPitagorica(s [][]int)` che riceve in input un valore di tipo `[][]int` nel parametro `s` e, come mostrato nell'**Esempio di esecuzione**, stampa la tavola pitagorica corrispondete ai valori memorizzati `s`.

##### Esempio d'esecuzione:

```text
$ go run tavola_pitagorica.go 5
   1    2    3    4    5 
   2    4    6    8   10 
   3    6    9   12   15 
   4    8   12   16   20 
   5   10   15   20   25 

$ go run tavola_pitagorica.go 10
   1    2    3    4    5    6    7    8    9   10 
   2    4    6    8   10   12   14   16   18   20 
   3    6    9   12   15   18   21   24   27   30 
   4    8   12   16   20   24   28   32   36   40 
   5   10   15   20   25   30   35   40   45   50 
   6   12   18   24   30   36   42   48   54   60 
   7   14   21   28   35   42   49   56   63   70 
   8   16   24   32   40   48   56   64   72   80 
   9   18   27   36   45   54   63   72   81   90 
  10   20   30   40   50   60   70   80   90  100
```
# Coppie

Scrivere un programma che:
* legga da **riga di comando** un numero intero `n`;
* utilizzi le funzioni `CreaSliceBidimensionale(l int) [][]int` e `AssegnaSliceBidimensionale(s [][]int) ` descritte nell'Esercizio 4 per inizializzare una variabile `s` di tipo `[][]int` con lunghezza/capacità pari a `n` in cui ogni elemento `s[i]` è un valore di tipo `[]int` con lunghezza/capacità pari a `n`;
* stampi a video tutte le coppie di indici `(i, j)`, tali che `s[i][j]` è uguale a `1`.

Oltre alle funzioni `main()`, `CreaSliceBidimensionale()`, e `AssegnaSliceBidimensionale()` devono essere definite ed utilizzate almeno le seguenti funzioni:
* una funzione `Coppie(s [][]int) (coppie [][]int)` che riceve in input un valore `[][]int` nel parametro `s` e restituisce il valore di tipo `[][]int` nella variabile `coppie` in cui sono memorizzate tutte le coppie di indici `(i, j)`, tali che `s[i][j]` è uguale a `1`.

##### Esempio d'esecuzione:

```text
$ go run coppie.go 3
la slice originale è:
        0       1       1
        1       1       0
        1       1       0
gli indirizzi degli elementi uguali a 1 sono:
[0 1]
[0 2]
[1 0]
[1 1]
[2 0]
[2 1]
```
