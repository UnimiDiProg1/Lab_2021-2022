package main

import "fmt"

func numeroDiGiorni(mese string) int{

	var numeroDiGiorni int

	switch mese{

		case "gennaio":
			numeroDiGiorni = 31
		case "febbraio":
			numeroDiGiorni = 28
		case "marzo":
			numeroDiGiorni = 31
		case "aprile":
			numeroDiGiorni = 30
		case "maggio":
			numeroDiGiorni = 31
		case "giugno":
			numeroDiGiorni = 30
		case "luglio":
			numeroDiGiorni = 31
		case "agosto":
			numeroDiGiorni = 31
		case "settembre":
			numeroDiGiorni = 30
		case "ottobre":
			numeroDiGiorni = 31
		case "novembre":
			numeroDiGiorni = 30
		case "dicembre":
			numeroDiGiorni = 31
		default:
			numeroDiGiorni = -1

	}

	return numeroDiGiorni
}

func main(){

	var mese string
	var numerodigiorni int

	fmt.Print("Inserire mese: ")
	fmt.Scan(&mese)

	numerodigiorni = numeroDiGiorni(mese)

	if numerodigiorni == -1{
		fmt.Println("Il nome del mese non è corretto.")
	} else{
		fmt.Println("Il numero di giorni di", mese, "è", numerodigiorni)
	}
}

