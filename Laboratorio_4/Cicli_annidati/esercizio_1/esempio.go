package main

import "fmt"

func main() {
	var n int = 3

	var i, j int
	for i = 0; i < n; i++ {
		for j = 0; j < n; j++ {
			fmt.Print("*")
		}
	}
	fmt.Println()
}
