package main

import (
	"fmt"
	"os"
)

func main() {
	
	a := os.Args[1]

	for i:=0; i<len(a)-1; i++ {
		if a[i] < a[i+1] {
			fmt.Print(string(a[i]))
		}
	}
	
	fmt.Println()
}
