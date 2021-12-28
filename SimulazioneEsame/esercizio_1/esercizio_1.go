package main

import(
"fmt"
"os"
"strconv"
"math"
"strings"
)

func maggioreDellaSoglia(valore float64) bool { //bonus
	soglia,_ := strconv.ParseFloat(os.Args[2],64)
	epsilon,_ := strconv.ParseFloat(os.Args[3],64)

	if valore > soglia + epsilon {
		return true
	}
	
	return false
}

func main(){

	estrai_a := strings.Split(os.Args[1],"x^2") // divide la stringa in a e quello che viene dopo "x^2"
	a,_ := strconv.ParseFloat(estrai_a[0],64)

	estrai_b := strings.Split(estrai_a[1],"x") // divide la stringa RIMANENTE in b e quello che viene dopo "x"
	b,_ := strconv.ParseFloat(estrai_b[0],64)

	estrai_c := strings.Split(estrai_b[1],"=")  // divide la stringa RIMANENTE in c e quello che viene dopo "="
	c,_ := strconv.ParseFloat(estrai_c[0],64)

	discriminante := b*b - (4*a*c)

	if discriminante < 0 {
		fmt.Println("Non ci sono soluzioni reali")
	} else if discriminante == 0{
		sol := -b/2*a
		fmt.Printf("Esiste un'unica soluzione reale: %.3f \n", sol)
		
		//bonus
		if maggioreDellaSoglia(sol) {
			fmt.Printf("La soluzione %.3f è maggiore della soglia\n", sol)
		}
		
	} else {
		sol1 := (-b+math.Sqrt(discriminante))/(2*a)
		sol2 := (-b-math.Sqrt(discriminante))/(2*a)
		fmt.Printf("Esistono due soluzioni reali: %.3f e %.3f\n", sol1, sol2)
		
		//bonus
		if maggioreDellaSoglia(sol1) {
			fmt.Printf("La soluzione %.3f è maggiore della soglia\n", sol1)
		}

		if maggioreDellaSoglia(sol2) {
			fmt.Printf("La soluzione %.3f è maggiore della soglia\n", sol2)
		}
	}
}