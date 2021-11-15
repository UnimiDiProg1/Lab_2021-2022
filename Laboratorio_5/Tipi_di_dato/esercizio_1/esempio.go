package main

import "fmt"

func main() {
	var x int = 10
	var y float64 = 2.5
	t := 100
	t, z, k := x/int(y), float64(x)/y, t*2
	fmt.Println(t, z, k)
}
