package main

import "fmt"

func main() {
	b := []int{1, 2, 3, 4, 5}
	stampa(b)

	b = modifica(b)
	stampa(b)

	b = eliminaUltimoElemento(b)
	stampa(b)
}

func stampa(sl []int) {
	for _, v := range sl {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func modifica(slIn []int) (slOut []int) {
	slOut = slIn
	for i := range slOut {
		slOut[i] *= 2
	}
	return
}

func eliminaUltimoElemento(slIn []int) (slOut []int) {
	slOut = slIn[:len(slIn)-1]
	return
}
