# Qual è l'output?

Qual è l'output di questo programma?

```go
package main

import "fmt"

func main() {
	var a [4]int = [4]int{1, 2, 3, 4}

	for i := 1; i <= len(a); i++ {
		fmt.Print(a[i], " ")
	}

}
```
# Qual è l'output?

Qual è l'output di questo programma?

```go
package main

import "fmt"

func main() {
	b := [3]rune{'a', 'b', 'c'}

	for i := range b {
		fmt.Println(i)
	}

}
```
# Qual è l'output?

Qual è l'output di questo programma?

```go
package main

import "fmt"

func main() {
	var a [5]string

	a[1] = "hello"
	a[4] = "world"

	for i := range a {
		fmt.Print(a[i])
	}

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

	for _, v := range a {
		v *= 2
	}

	fmt.Println(a)
}
```
# Trova l'errore

Il seguente programma non può essere eseguito a causa di un errore. Qual è l'errore? Come è possibile risolvere il problema?

```go
package main

func main() {
	var n int = 4

	var a [n]int

	for i := 0; i < n; i++ {
		a[i] = i
	}
}
```
# Qual è l'output?

Qual è l'output di questo programma?

```go
package main

import "fmt"

func main() {
	var a = [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(a)
	modifica(a)
	fmt.Println(a)
}

func modifica(a [6]int) {
	for i := range a {
		a[i] *= 2
	}
}
```
