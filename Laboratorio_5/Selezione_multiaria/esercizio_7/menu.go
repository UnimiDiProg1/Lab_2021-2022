package main

import "fmt"

func main() {
	var o int
	var tot int
	var n int

	for {
		fmt.Println(`Cosa vuoi ordinare?
		1. patatine 2€
		2. hamburger 5€
		3. cocacola 2€
		0. termina`)
			
		fmt.Scan(&o)
		switch o {
		case 1:
			fmt.Println("patatine? ottimo, quante?")
			fmt.Scan(&n)
			tot = tot + (n * 2)
		case 2:
			fmt.Println("hamburger? ottimo, quanti?")
			fmt.Scan(&n)
			tot = tot + (n * 5)
		case 3:
			fmt.Println("cocacola? ottimo, quante?")
			fmt.Scan(&n)
			tot = tot + (n * 2)
		case 0:
			fmt.Println("Sono", tot, "euro + 2 di consegna. Totale:", (tot + 2))
			return
		}
	}

}
