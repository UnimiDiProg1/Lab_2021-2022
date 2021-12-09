package main

import (
"fmt"
)

type Indirizzo struct{

        via     string
        CAP     string
        città   string
        numeroCivico uint

}

type Contatto struct{

        cognome string
        nome string
        telefono string
        indirizzo Indirizzo

}

type Rubrica []Contatto


func NuovaRubrica() Rubrica{

        return Rubrica{}

}

func InserisciContatto(r Rubrica, c Contatto) Rubrica{

        return append(r,c)

}
func main(){

        rubrica := NuovaRubrica()

        indirizzo := Indirizzo{"via x", "20138", "Milano", 5}
        contatto := Contatto{"nome", "cognome", "212312", indirizzo}

        rubrica = InserisciContatto(rubrica, contatto)

        indirizzo2 := Indirizzo{"via y", "20099", "Sesto San Giovanni", 8}
        contatto2 := Contatto{"altronome", "altrocognome", "4484848", indirizzo2}

        rubrica = InserisciContatto(rubrica, contatto2)

        fmt.Println(rubrica)

}