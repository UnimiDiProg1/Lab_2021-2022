package main

import "fmt"
import "os"
import "strconv"

func main() {

	prodotto := 1
	for _, v := range os.Args[1:] {

		if n, err := strconv.Atoi(v); err == nil {
			prodotto *= n
		}

	}

	fmt.Println("Il risultato della moltiplicazione tra i numeri interi è", prodotto)

}
