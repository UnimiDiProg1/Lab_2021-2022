package main

import "fmt"

func main() {

	var n int

	fmt.Scan(&n)

	accumulatore := 1

	for i := 1; i <= n; i++ {

		switch {
		case i < 5:
			accumulatore *= i
		case i < 10:
			accumulatore += i
		}

	}

	fmt.Println("Accumulatore =", accumulatore)
}
