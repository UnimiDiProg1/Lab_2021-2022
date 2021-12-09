package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
	"sort"
)

func main() {

	StampaOccorrenze(Occorrenze(LeggiTesto()))

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

func StampaOccorrenze(occorrenze map[rune]int) {

	chiavi := []string{}
	for k := range occorrenze {
		chiavi = append(chiavi, string(k))
	}

	sort.Strings(chiavi)

	fmt.Println("Occorrenze:")
	for _, k := range chiavi {
		var runes []rune = []rune(k)
		var r rune = runes[0]
		fmt.Printf("%c: %s\n", r, strings.Repeat("*", occorrenze[r]))
	}
}


