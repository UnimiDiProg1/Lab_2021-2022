package main

import "fmt"

func main() {

	var (
		m, q,
		x, y float64
	)

	fmt.Print("Inserisci m e q: ")
	fmt.Scan(&m, &q)

	fmt.Print("Inserisci x e y: ")
	fmt.Scan(&x, &y)

	if y > m*x+q {
		fmt.Println("Il punto sta sopra la retta")
	} else if y < m*x+q {
		fmt.Println("Il punto sta sotto la retta")
	} else {
		fmt.Println("Il punto appartiene alla retta")
	}
}
