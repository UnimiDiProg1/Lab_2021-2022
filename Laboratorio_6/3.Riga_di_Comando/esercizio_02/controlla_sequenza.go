package main

import (
	"os"
	"strconv"
)
import "fmt"

func main() {
	sequenzaValida, numeri, posizione := LeggiNumeri()

	if sequenzaValida {
		sequenzaValida, posizione = ControllaNumeri(numeri)
		if sequenzaValida {
			fmt.Println("Sequenza valida.")
		} else {
			fmt.Printf("Valore in posizione %d non valido.\n", posizione)
		}
	} else {
		fmt.Printf("Valore in posizione %d non valido.\n", posizione)
	}
}

func LeggiNumeri() (sequenzaValida bool, numeri []int, posizione int) {
	sequenzaValida = true
	posizione = -1

	for i, valore := range os.Args[1:] {
		if n, err := strconv.Atoi(valore); err == nil {
			numeri = append(numeri, n)
		} else {
			sequenzaValida = false
			posizione = i
			numeri = []int{}
			break
		}
	}
	return
}

func ControllaNumeri(numeri []int) (sequenzaValida bool, posizione int) {
	sequenzaValida = true
	posizione = -1

	for i := 1; i < len(numeri); i++ {

		if i%2 == 0 {
			if numeri[i] <= numeri[i-1] {
				sequenzaValida = false
			}
		} else {
			if numeri[i] >= numeri[i-1] {
				sequenzaValida = false
			}
		}

		if !sequenzaValida {
			posizione = i
			break
		}
	}

	return

}
