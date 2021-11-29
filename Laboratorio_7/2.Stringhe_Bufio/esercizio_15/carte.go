package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {

	rand.Seed(int64(time.Now().Nanosecond()))

	n := LeggiNumero()
	EstraiCarte(GeneraMazzo(), n)
}

func LeggiNumero() int {
	n, _ := strconv.Atoi(os.Args[1])
	return n
}

func EstraiCarta(mazzo string) (cartaEstratta rune, mazzoResiduo string) {
	mazzoRune := []rune(mazzo)
	indice := rand.Intn(len(mazzoRune))
	cartaEstratta = mazzoRune[indice]
	mazzoRune = append(mazzoRune[:indice], mazzoRune[indice+1:]...)
	mazzoResiduo = string(mazzoRune)
	return
}

func GeneraMazzo() (mazzo string) {
	const asso = '\U0001F0B1'
	for i := 0; i < 10; i++ {
		mazzo += string(asso + i)
	}
	return
}

func EstraiCarte(mazzo string, iterazioni int) {
	for i := 0; i < iterazioni; i++ {
		var carta rune
		carta, mazzo = EstraiCarta(mazzo)
		fmt.Printf("Estratta la carta %c - Carte rimaste nel mazzo: %s\n", carta, mazzo)
	}
}
