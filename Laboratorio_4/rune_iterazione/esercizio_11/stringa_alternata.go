package main

import (
	"fmt"
	"strings"
)

func main() {

	var s1, s2 string
	fmt.Scan(&s1, &s2)

	fmt.Println(StringheAlternate(s1, s2))

}

func StringheAlternate(s1, s2 string) (risultato string) {

	lunghezzaMassima := len(s1)
	if lunghezzaMassima < len(s2) {
		lunghezzaMassima = len(s2)
	}

	s1 += strings.Repeat("-", lunghezzaMassima-len(s1))
	s2 += strings.Repeat("-", lunghezzaMassima-len(s2))

	for i := range s1 {
		risultato += string(s1[i]) + string(s2[i])
	}

	return
}
