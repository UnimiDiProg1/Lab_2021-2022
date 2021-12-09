package main

import "fmt"

func main() {

	rubrica := NuovaRubrica()

	rubrica = InserisciContatto(rubrica, "Mario", "Rossi",
		"Via Festa del Perdono", 11, "20122", "Milano", "02503111")
	rubrica = InserisciContatto(rubrica, "Mario", "Rossi",
		"Via Festa del Perdono", 11, "", "", "")
	rubrica = InserisciContatto(rubrica, "Anna", "Rossi",
		"Via Festa del Perdono", 11, "20122", "Milano", "02503111")
	rubrica = InserisciContatto(rubrica, "Carlo", "Rossi",
		"Via Festa del Perdono", 11, "20122", "Milano", "02503111")

	StampaRubrica(rubrica)

	rubrica = EliminaContatto(rubrica, "Mario", "Rossi")
	rubrica = EliminaContatto(rubrica, "Carlo", "Verdi")

	StampaRubrica(rubrica)

	rubrica = AggiornaContatto(rubrica, "Anna", "Rossi", "Via S. Sofia", 25, "20122", "Milano", "")

	StampaRubrica(rubrica)
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

type Rubrica []Contatto

func NuovaRubrica() Rubrica {
	return Rubrica{}
}

func StampaContatto(contatto Contatto) {
	fmt.Printf("%s %s: %s %d, %s, %s - Tel. %s\n", contatto.nome,
		contatto.cognome, contatto.indirizzo.via, contatto.indirizzo.numeroCivico, contatto.indirizzo.città,
		contatto.indirizzo.cap, contatto.telefono)
}

func StampaRubrica(rubrica Rubrica) {
	fmt.Println("Rubrica:")
	for i, c := range rubrica {
		fmt.Printf("[%d] - ", i+1)
		StampaContatto(c)
	}
}

func InserisciContatto(rubrica Rubrica, nome, cognome, via string, numero uint, cap, città, telefono string) Rubrica {
	for _, c := range rubrica {
		if c.nome == nome && c.cognome == cognome {
			return rubrica
		}
	}
	return append(rubrica, Contatto{nome, cognome,
		Indirizzo{via, numero, città, cap},
		telefono})
}

func EliminaContatto(rubrica Rubrica, nome, cognome string) Rubrica {
	for i, contatto := range rubrica {
		if contatto.nome == nome && contatto.cognome == cognome {
			rubrica = append(rubrica[:i], rubrica[i+1:]...)
		}
	}
	return rubrica
}

func AggiornaContatto(rubrica Rubrica, nome, cognome, via string, numero uint, cap, città, telefono string) Rubrica {
	for i, contatto := range rubrica {
		if contatto.nome == nome && contatto.cognome == cognome {
			rubrica[i].telefono = telefono
			rubrica[i].indirizzo = Indirizzo{via, numero, cap, città}
		}
	}
	return rubrica
}
