package main

import "fmt"

func main() {
	b := [3]rune{'a', 'b', 'c'}

	for i := range b {
		fmt.Println(i)
	}

}
