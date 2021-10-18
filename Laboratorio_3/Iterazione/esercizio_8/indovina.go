package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(int64(time.Now().Nanosecond()))

	var numeroEstratto int = rand.Intn(100) + 1
	var tentativi int = 1

	for {
		var numeroLetto int
		fmt.Print("Tentativo n° ", tentativi, ": ")
		fmt.Scan(&numeroLetto)

		if numeroLetto == numeroEstratto {
			break
		} else if numeroLetto > numeroEstratto {
			fmt.Println("Troppo alto! Riprova!")
		} else {
			fmt.Println("Troppo basso! Riprova!")
		}

		tentativi++
	}

	fmt.Println("Hai indovinato in ", tentativi, " tentativi!")
}
