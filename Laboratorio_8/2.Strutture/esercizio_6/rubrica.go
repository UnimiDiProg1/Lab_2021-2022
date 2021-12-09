package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	listaComandi := LeggiTesto()
	rubrica := NuovaRubrica()

	for _, comando := range listaComandi {
		rubrica = InterpretaComando(comando, rubrica)
	}
}

func LeggiTesto() (testo []string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		testo = append(testo, scanner.Text())
	}
	return
}

func Sottostringhe(s string) (sottostringhe []string) {
	return strings.Split(s, ";")
}

func InterpretaComando(comando string, rubrica Rubrica) Rubrica {
	tokenComando := Sottostringhe(comando)
	switch tokenComando[0] {
	case "I":
		numero, _ := strconv.Atoi(tokenComando[4])
		rubrica = InserisciContatto(rubrica, tokenComando[1], tokenComando[2], tokenComando[3], uint(numero), tokenComando[5], tokenComando[6], tokenComando[7])
	case "E":
		rubrica = EliminaContatto(rubrica, tokenComando[1], tokenComando[2])
	case "S":
		StampaRubrica(rubrica)
	case "A":
		numero, _ := strconv.Atoi(tokenComando[4])
		rubrica = AggiornaContatto(rubrica, tokenComando[1], tokenComando[2], tokenComando[3], uint(numero), tokenComando[5], tokenComando[6], tokenComando[7])
	}
	return rubrica
}

type Indirizzo struct {
	via          string
	numeroCivico uint
	cap, città   string
}

type Contatto struct {
	nome, cognome string
	indirizzo     Indirizzo
	telefono      string
}

type Rubrica map[string]Contatto

func NuovaRubrica() Rubrica {
	return make(Rubrica)
}

func StampaContatto(contatto Contatto) {
	fmt.Printf("%s %s: %s %d, %s, %s - Tel. %s\n", contatto.nome,
		contatto.cognome, contatto.indirizzo.via, contatto.indirizzo.numeroCivico, contatto.indirizzo.città,
		contatto.indirizzo.cap, contatto.telefono)
}

func StampaRubrica(rubrica Rubrica) {
	fmt.Println("Rubrica:")
	var i int
	for _, c := range rubrica {
		fmt.Printf("[%d] - ", i+1)
		StampaContatto(c)
		i++
	}
}

func InserisciContatto(rubrica Rubrica, nome, cognome, via string, numero uint, cap, città, telefono string) Rubrica {
	chiave := nome + cognome
	_, ok := rubrica[chiave]
	if ok {
		return rubrica
	}
	rubrica[chiave] = Contatto{nome, cognome,
		Indirizzo{via, numero, città, cap},
		telefono}
	return rubrica
}

func EliminaContatto(rubrica Rubrica, nome, cognome string) Rubrica {
	chiave := nome + cognome
	_, ok := rubrica[chiave]
	if ok {
		delete(rubrica, chiave)
	}
	return rubrica
}

func AggiornaContatto(rubrica Rubrica, nome, cognome, via string, numero uint, cap, città, telefono string) Rubrica {
	chiave := nome + cognome
	_, ok := rubrica[chiave]
	if !ok {
		return rubrica
	}
	rubrica[chiave] = Contatto{nome, cognome,
		Indirizzo{via, numero, città, cap},
		telefono}
	return rubrica
}
