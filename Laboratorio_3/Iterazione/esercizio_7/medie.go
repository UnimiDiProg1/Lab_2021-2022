package main

import (
	"fmt"
)

func main() {
	var numeriLetti int = 0
	var mAritmetica float64 = 0.0

	fmt.Print("Inserisci una sequenza di numeri (interrompi con numero<=0): ")
	for {
		var x float64
		fmt.Scan(&x)

		if x > 0 {
			numeriLetti++
			mAritmetica += x
		} else {
			break
		}
	}

	mAritmetica /= float64(numeriLetti)

	fmt.Println("Media aritmetica:", mAritmetica)

}
