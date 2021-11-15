package main

import "fmt"

func main() {
	var operando1, operando2 float64
	var operatore rune
	var tot float64
	var flag bool = true

	fmt.Println("inserisci un'operazione aritmetica")
	fmt.Scanf("%f %c %f", &operando1, &operatore, &operando2)

	switch operatore {
		case '+':
			tot=operando1+operando2
		case '-':
			tot=operando1-operando2
		case '/':
			tot=operando1/operando2
		case '*':
			tot=operando1*operando2
		default:
			flag = false
	}
	
	if flag {
		fmt.Println("Risultato:", tot)
	} else {
			fmt.Println("Operatore non riconosciuto")
	}
}
