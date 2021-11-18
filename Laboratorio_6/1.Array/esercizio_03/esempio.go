package main

import "fmt"

func main() {
	var a [5]string

	a[1] = "hello"
	a[4] = "world"

	for i := range a {
		fmt.Print(a[i])
	}

}
