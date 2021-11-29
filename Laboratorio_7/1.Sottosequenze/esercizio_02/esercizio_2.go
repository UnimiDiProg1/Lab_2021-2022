package main

import (
	"fmt"
	"os"
)

func isBalanced(sequence string) bool {
	var counter int
	for _,v:=range(sequence) {
		if v == '(' {
			counter++
		} else if v == ')' {
			counter--
			
			if counter < 0 {
				return false
			}
		}
	}
	
	if counter == 0 {
		return true
	}
	
	return false
}

func main() {
	var s string
	
	if len(os.Args)!=2{

		fmt.Println("Richiesto un solo argomento.")
		os.Exit(1)

	}

	s = os.Args[1]
	
	for _,v:=range(s) {
		if v != '(' && v != ')' {
			fmt.Println("L'input fornito non aveva esclusivamente parentesi tonde.")
			return
		}
	}
	
	if isBalanced(s) {
		fmt.Println("bilanciata")
	} else {
		fmt.Println("non bilanciata")
	}
	
	fmt.Println("---")

	fmt.Println("Sottosequenze bilanciate:")

	counter := 0

	for i := 0; i < len(s); i++ {
		for j := i+1; j <= len(s); j++ {
			subsequence := s[i:j]
			if isBalanced(subsequence) {
				counter++
				fmt.Print(counter,") ",subsequence, "\n")
			}
		}
	}

	if(counter==0){
		fmt.Println("Nessuna.")
	}
	
}