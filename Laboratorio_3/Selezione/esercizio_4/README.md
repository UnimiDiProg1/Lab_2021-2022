# Qual è l'output?

Confronta i programmi che seguono. Cosa stampano nel caso in cui l'utente inserisca da **standard input** il valore `150`? E se invece inserisse `40`? I due programmi funzionano allo stesso modo?

```go
package main

import "fmt"

func main() {

	var a int

	fmt.Scan(&a)

	if a < 100 {
		fmt.Println("a minore di 100")
	} else if a < 200 {
		fmt.Println("a compreso tra 100 e 200")
	} else {
		fmt.Println("a maggiore o uguale a 200")
	}

}
```

```go
package main

import "fmt"

func main() {

	var a int

	fmt.Scan(&a)

	if a < 100 {
		fmt.Println("a minore di 100")
	} 
	if a < 200 {
		fmt.Println("a compreso tra 100 e 200")
	} else {
		fmt.Println("a maggiore o uguale a 200")
	}

}
```
