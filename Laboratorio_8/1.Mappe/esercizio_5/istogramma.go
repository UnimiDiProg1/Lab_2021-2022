package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {

	occorrenze := Occorrenze(LeggiTesto())

	fmt.Println("Occorrenze:")
	for k, v := range occorrenze {
		fmt.Printf("%c: %s\n", k, strings.Repeat("*", v))
	}
}

func LeggiTesto() string {
	scanner := bufio.NewScanner(os.Stdin)
	testo := ""
	for scanner.Scan() {
		testo += scanner.Text()
	}
	return testo
}

func Occorrenze(s string) map[rune]int {
	occorrenze := map[rune]int{}
	for _, l := range s {
		if unicode.IsLetter(l) {
			occorrenze[l]++
		}
	}
	return occorrenze
}
