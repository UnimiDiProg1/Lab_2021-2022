package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	sequenza := os.Args[1:]

	sottosequenze := GeneraSottosequenze(sequenza)

	StampaSottosequenze(sottosequenze)
}

func GeneraSottosequenze(sequenza []string) (sottosequenze map[string]int) {
	sottosequenze = make(map[string]int)
	for i, c1 := range sequenza {
		for j, c2 := range sequenza[i+1:] {
			if c1 == c2 {
				sottosequenze[strings.Join(sequenza[i:i+j+2], " ")]++
			}
		}
	}
	return
}

func StampaSottosequenze(sottosequenze map[string]int) {

	for sequenza, occorrenze := range sottosequenze {
	
		if sequenza[0]==sequenza[len(sequenza)-1] && len(sequenza)>=5 {

			fmt.Printf("%s -> Occorrenze: %d\n", sequenza, occorrenze)

		}
	}
		
}
