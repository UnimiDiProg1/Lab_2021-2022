package main

import "fmt"

func main() {

	s := "Ciao René!"
	numerocaratteri := 0
	numeroiterazione := 0
	for i, c := range s {
		fmt.Print("Iterazione ", numeroiterazione, ": In posizione ", i,
			" inizia la sottosequenza di byte che codifica il carattere ",
			string(c), "\n")
		numerocaratteri++
		numeroiterazione++
	}
	fmt.Println("Numero iterazioni:", numeroiterazione)
	fmt.Println("Numero di byte utilizzati per rappresentare la stringa:", len(s))
	fmt.Println("Numero di caratteri che definiscono la stringa:", numerocaratteri)
}
