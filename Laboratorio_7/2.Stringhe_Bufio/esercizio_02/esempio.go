package main

import "fmt"

func main() {

	//strFrom := "ABDEF"
	//strFrom[2] = 'C' /* error: "cannot assign to strFrom[2]" */

	/* (1) */
	strFrom1 := "ABÈEF"
	// len(sliceDiRune) è pari al numero di caratteri Unicode presenti in strFrom1
	sliceDiRune := []rune(strFrom1)
	sliceDiRune[2] = 'C'
	strTo1 := string(sliceDiRune)
	fmt.Println(strTo1)
	/* (1) - end */

	/* (2) */
	strFrom2 := "ABÈEF"
	//strTo2 := strFrom2[:2]+string('C')+strFrom2[3:]
	strTo2 := strFrom2[:2] + "C" + strFrom2[3:]
	fmt.Println(strTo2)
	/* (2) - end */

}
