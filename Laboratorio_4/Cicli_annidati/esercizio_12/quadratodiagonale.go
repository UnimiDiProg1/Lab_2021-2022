package main

import (
	"fmt"
)

func main() {

	var numero int
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&numero)

	var i, j int
	for i = 0; i < numero; i++ {
		for j = 0; j < numero; j++ {
			if i == j {
				fmt.Print("o ")
			} else if i < j {
				fmt.Print("+ ")
			} else {
				fmt.Print("* ")
			}
		}
		fmt.Println()
	}

}
