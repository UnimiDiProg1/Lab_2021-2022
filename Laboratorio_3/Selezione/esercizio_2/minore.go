package main

import "fmt"

func main() {

	var a, b, c int
	var m int

	fmt.Scan(&a, &b, &c)

	if a < b {
		if a < c {
			m = a
		} else {
			m = c
		}
	} else {
		if b < c {
			m = b
		} else {
			m = c
		}
	}

	fmt.Println(m)

}
