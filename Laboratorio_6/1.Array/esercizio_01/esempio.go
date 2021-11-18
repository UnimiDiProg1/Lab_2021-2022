package main

import "fmt"

func main() {
	var a [4]int = [4]int{1, 2, 3, 4}

	for i := 1; i <= len(a); i++ {
		fmt.Print(a[i], " ")
	}

}
