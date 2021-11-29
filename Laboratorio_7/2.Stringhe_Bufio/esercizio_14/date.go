package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sort"
)

func main() {
	date := LeggiTesto()

	var dateFormattate []string
	for _, data := range date {
		dateFormattate = append(dateFormattate, Formatta(data))
	}

	sort.Strings(dateFormattate)
	for _, data := range dateFormattate {
		fmt.Println(data)
	}

}

func LeggiTesto() (date []string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		riga := scanner.Text()
		if riga == "" {
			return
		}
		date = append(date, riga)
	}
	return
}


func Formatta(data string) string{

	//separiamo la data in giorno mese e anno usando il carattere separatore /
	dataScomposta := strings.Split(data,"/")

	var giorno, mese, anno string

	var zeroG, zeroM string

	//se il primo campo ha due o meno caratteri, vuol dire che è il giorno
	if len(dataScomposta[0])<=2{

		giorno = dataScomposta[0]
		mese = dataScomposta[1]
		anno = dataScomposta[2]

	} else{ //altrimenti è l'anno

		giorno = dataScomposta[2]
		mese = dataScomposta[1]
		anno = dataScomposta[0]

	}

	//quando il giorno o il mese hanno un solo carattere, devo aggiungere uno zero davanti
	//di default zeroG e zeroM sono stringhe vuote, le valorizzo a "0" solo se serve
	if len(giorno)==1{
		zeroG = "0"
	}

	if len(mese)==1{
		zeroM = "0"
	}


	return anno+"-"+zeroM+mese+"-"+zeroG+giorno

}
