package main

import "fmt"

func main() {

	var a int

	fmt.Scan(&a)

	if a < 100 {
		fmt.Println("a minore di 100")
	}
	if a < 200 {
		fmt.Println("a compreso tra 100 e 200")
	} else {
		fmt.Println("a maggiore o uguale a 200")
	}

}
