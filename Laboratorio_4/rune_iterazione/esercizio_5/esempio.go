package main

import "fmt"

func main() {

	s := "Ciao, come va?"
	// s è interamente definita da caratteri considerati nello standard US-ASCII

	for i := 0; i < len(s); i++ {
		fmt.Print(string(s[i]))
	}
	fmt.Println()

	s = "Ciao, come è andata?"
	// s non è interamente definita da caratteri considerati nello standard US-ASCII

	for i := 0; i < len(s); i++ {
		fmt.Print(string(s[i]))
	}
	fmt.Println()

}
