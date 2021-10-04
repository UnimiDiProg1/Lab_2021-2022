package main

import "fmt" // il package fmt va tra virgolette

func main() { // il main necessita delle parentesi tonde

	var x int 
    x = 10
	var y int
    y = 15

    var sum int // la variabile sum non è stata dichiarata
	sum = x + y

	fmt.Println("La somma è ", sum) // la parentesi tonda non era chiusa

}
// } una parentesi graffa di troppo
