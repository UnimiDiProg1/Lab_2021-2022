package main

import "fmt"

func main() {

	var n int

	fmt.Scan(&n)

	var somma int

	for i := 0; i <= n; i++ {

		switch i % 2 {
		case 0:
			fmt.Println(i, "pari!")
		case 1:
			break
		}

		somma += i
	}

	fmt.Println("Somma =", somma)
}
