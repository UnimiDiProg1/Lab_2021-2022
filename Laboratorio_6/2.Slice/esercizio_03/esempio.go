package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7}

	fmt.Printf("a - %T: %v\n", a, a)

	sl1 := a[:] // slicing
	sl2 := sl1[1:3]

	fmt.Printf("len(sl1) = %d, cap(sl1) = %d\n", len(sl1), cap(sl1))
	fmt.Printf("sl1 - %T: %v\n", sl1, sl1)
	fmt.Printf("len(sl2) = %d, cap(sl2) = %d\n", len(sl2), cap(sl2))
	fmt.Printf("sl2 - %T: %v\n", sl2, sl2)

	sl2 = sl2[:len(sl2)+1] // reslicing
	fmt.Printf("\nsl2 - %T: %v\n", sl2, sl2)

	sl2 = sl2[:cap(sl2)]
	fmt.Printf("sl2 - %T: %v\n", sl2, sl2)

	elem, sl1 := sl1[0], sl1[1:]
	fmt.Printf("\nelem - %T: %v\n", elem, elem)
	fmt.Printf("sl1 - %T: %v\n", sl1, sl1)

	/* una slice s non può essere modificata per accedere ad elementi
	   dell'array (a cui si riferisce) che precedono quello contenuto in
	   s[0]; l’istruzione s = s[-1:] genera un errore */

}
