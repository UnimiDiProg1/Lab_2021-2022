package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Inserisci testo (termina con CTRL+D):")
	fmt.Print("Testo letto:\n", LeggiTesto())

}

func LeggiTesto() (testo string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		testo += scanner.Text() + "\n"
	}
	return
}
