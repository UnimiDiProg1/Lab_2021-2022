package main

import "fmt"

func main() {

	s := "ciao, come va?"
	/* s è interamente definita da caratteri considerati nello standard US-ASCII */

	fmt.Println(string(s[0]) + string(s[len(s)-2]) + string(s[len(s)-1]))
}
