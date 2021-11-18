package main

import "fmt"

func main(){

	var n int
	fmt.Scan(&n)

	var slice []int
	slice = make([]int, n)

	fmt.Println("Inserisci numeri:")

	for i:=0; i<n; i++{

		fmt.Scan(&slice[i])

	}

	fmt.Println("Ordine inverso:")

	for i:=len(slice)-1; i>=0; i--{

		fmt.Print(slice[i], " ")

	}



}