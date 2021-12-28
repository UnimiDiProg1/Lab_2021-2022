package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Utenza struct {
	telefono, sim string
}

type RegistroTelefonico map[string][]string

func main() {
	utenze := LeggiUtenze()
	registro := InizializzaRegistro()

	for _, utenza := range utenze {
		AggiungiUtenza(registro, utenza)
	}

	for telefono, _ := range registro {
		if telefono[:3]=="338"{
			simRecente := Identifica(registro,telefono)
			fmt.Println("Il numero",telefono,"è associato alla sim",simRecente)
		}
	}
}

func LeggiUtenze() (utenze []Utenza) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		riga := strings.Split(scanner.Text(), ";")
		utenze = append(utenze, Utenza{riga[0], riga[1]})
	}
	return
}


func InizializzaRegistro() (registro RegistroTelefonico) {
	registro = make(RegistroTelefonico)
	return
}

func AggiungiUtenza(registro RegistroTelefonico, utenza Utenza) RegistroTelefonico {
	registro[utenza.telefono] = append(registro[utenza.telefono], utenza.sim)
	return registro
}

func Identifica(registro RegistroTelefonico, telefono string) (codiceSIM string) {
	listaSim := registro[telefono]
	l := len(listaSim)

	if l > 0 {
		codiceSIM = listaSim[l-1]
	}

	return
}