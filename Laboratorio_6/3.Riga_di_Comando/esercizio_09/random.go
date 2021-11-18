package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	soglia, _ := strconv.Atoi(os.Args[1])

	numeriCasuali := Genera(soglia)

	fmt.Println("Valori generati", numeriCasuali)
	fmt.Println("Valori sopra soglia:", numeriCasuali[:len(numeriCasuali)-1])
}

func Genera(soglia int) (numeriCasuali []int) {
	rand.Seed(int64(time.Now().Nanosecond()))

	for n := int(soglia) + 1; n > int(soglia); {
		n = rand.Intn(100)
		numeriCasuali = append(numeriCasuali, n)
	}

	return
}
