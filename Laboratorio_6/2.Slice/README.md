# Qual è l'output?

Qual è l'output di questo programma?

```go
package main

import "fmt"

func main() {

	var n int = 5

	var s []int
	s = make([]int, n)

	for i := 0; i < n; i++ {
		s[i] = i
	}

	fmt.Println(s)
}
```
# Qual è l'output?

Analizziamo l'output del seguente programma.

```go
package main

import "fmt"

func main() {
	a := [...]string{5:"E", 3:"C"} 
	// equivalente a: var a [6]string = [6]string{"", "", "", "C", "", "E"}

	fmt.Printf("Tipo: %T  - Valore: %v\n", a, a)

	for i := range a {
		fmt.Println("Indice", len(a)-1-i, " - Valore:", a[len(a)-1-i])
	}
	fmt.Println()

	for i:=len(a)-1; i>=0; i-- {
		fmt.Println("Indice", i, " - Valore:", a[i])
	}
	fmt.Println()

	for i:=0; i<len(a); i++ {
		fmt.Println("Indice", i, " - Valore:", a[i])
	}
	fmt.Println()
}
```
# Qual è l'output?

Analizziamo l'output del seguente programma.

```go
package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7}

	fmt.Printf("a - %T: %v\n", a, a)

	sl1 := a[:] // slicing
	sl2 := sl1[1:3]

	fmt.Printf("len(sl1) = %d, cap(sl1) = %d\n", len(sl1), cap(sl1))
	fmt.Printf("sl1 - %T: %v\n", sl1, sl1)
	fmt.Printf("len(sl2) = %d, cap(sl2) = %d\n", len(sl2), cap(sl2))
	fmt.Printf("sl2 - %T: %v\n", sl2, sl2)

	sl2 = sl2[:len(sl2)+1]   // reslicing
	fmt.Printf("\nsl2 - %T: %v\n", sl2, sl2)

	sl2 = sl2[:cap(sl2)]
	fmt.Printf("sl2 - %T: %v\n", sl2, sl2)

	elem, sl1 := sl1[0], sl1[1:]
	fmt.Printf("\nelem - %T: %v\n", elem, elem)
	fmt.Printf("sl1 - %T: %v\n", sl1, sl1)

	/* una slice s non può essere modificata per accedere ad elementi
	   dell'array (a cui si riferisce) che precedono quello contenuto in
	   s[0]; l’istruzione s = s[-1:] genera un errore */

}
```
# Qual è l'output?

Qual è l'output di questo programma?

```go
package main

import "fmt"

func main() {
	var a [6]int

	for i := range a {
		a[i] = i
	}

	var b []int
	b = a[:]

	for i := range b {
		b[i] = i * 2
	}

	fmt.Println(a)
}
```
# Qual è l'output?

Qual è l'output di questo programma?

```go
package main

import "fmt"

func main() {
	var a []int
	a = []int{0, 1, 2, 3, 4, 5, 6}

	var b []int
	b = a[2:4]

	b[0] = a[0]
	b[len(b)-1] = a[0]

	fmt.Println(a)
}
```
# Qual è l'output?

Qual è l'output di questo programma?

```go
package main

import "fmt"

func main() {
	var a = [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(a[:])

	b := a[:]
	b = modifica(b)

	fmt.Println(b)
}

func modifica(slIn []int) (slOut []int) {
    slOut = slIn
	for i := range slOut {
		slOut[i] *= 2
	}
	return
}
```
# Qual è l'output?

Qual è l'output di questo programma?

```go
package main

import "fmt"

func main() {
	b := []int{1,2,3,4,5}
	stampa(b)

	modifica(b)
	stampa(b)

	eliminaUltimoElemento(b)
	stampa(b)
}

func stampa(sl []int) {
	for _, v := range sl {
		fmt.Print(v, " ")
	}
	fmt.Println()
}
func modifica(sl []int) {
	for i := range sl {
		sl[i] *= 2
	}
}

func eliminaUltimoElemento(sl []int) {
	sl = sl[:len(sl)-1]
}
```
# Stampa in ordine inverso


Scrivere un programma che, dopo aver letto da **standard input** un numero intero `n`, chiede all'utente di inserire `n` numeri interi (sempre da **standard input**). 

Il programma deve stampare gli `n` numeri interi in ordine inverso rispetto a quello di inserimento.

*Suggerimento:* Per creare dinamicamente una slice si utilizzi la funzione `make()`.

##### Esempio d'esecuzione:

```text
$ go run stampa_rovescio.go 
9
Inserisci 9 numeri:
1 -12 3 -4 5 -6 7 -7 9
Numeri in ordine inverso:
9 -7 7 -6 5 -4 3 -12 1 

$ go run stampa_rovescio.go 
5
Inserisci 5 numeri:
1 2 3 4 5
Numeri in ordine inverso:
5 4 3 2 1 
```
# Qual è l'output?

Analizziamo l'output del seguente programma.

```go
package main

import (
	"fmt"
)

func main() {

	s1 := make([]int, 4, 8)

	for i := 0; i < len(s1); i++ {
		s1[i] = i+1
	}

	fmt.Println("1)\ns1:")
	stampa(s1)

	s2 := append(s1[2:], []int{10,20}...)
	fmt.Println("\n2)\ns2:")
	stampa(s2)
	fmt.Println("s1:")
	stampa(s1)

	s2[0] = -1
	s2 = append(s2, 100)
	fmt.Println("\n3)\ns2:")
	stampa(s2)
	fmt.Println("s1:")
	stampa(s1)

	s2 = append(s2, 1000, 2000)
	fmt.Println("\n4)\ns2:")
	stampa(s2)
	fmt.Println("s1:")
	stampa(s1)

	s2 = append(s2[:2], s2[len(s2)-2:]...)
	fmt.Println("\n5)\ns2:")
	stampa(s2)

}

func stampa(s []int) {
	var lunghezza int
	fmt.Println("a) ", s, len(s), cap(s))
	lunghezza = len(s)
	s = s[:cap(s)]
	fmt.Println("b) ", s, len(s), cap(s))
	s = s[:lunghezza]
}
```
# Qual è l'output?

Qual è l'output del seguente programma?

```go
package main

import (
	"fmt"
)

const Dimensione = 10

func main() {

	var a []int

	for i := 0; i < Dimensione; i++ {
		a = append([]int{i + 1}, a...)
	}

	fmt.Println(a)

	b := make([]int, Dimensione)

	copy(b, a[Dimensione/2:])

	fmt.Println(b)

}
```
