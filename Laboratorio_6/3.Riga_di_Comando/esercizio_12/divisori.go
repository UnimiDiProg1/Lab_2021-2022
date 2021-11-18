package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	n, _ := strconv.Atoi(os.Args[1])
	min, _ := strconv.Atoi(os.Args[2])
	max, _ := strconv.Atoi(os.Args[3])

	for a := 1; a <= n; a++ {
		for b := a; b <= n; b++ {
			divisoriComuni := DivisoriComuni(DivisoriPropri(a), DivisoriPropri(b))
			if (ÈPerfetto(a) || ÈPerfetto(b)) && divisoriComuni >= min && divisoriComuni <= max {
				fmt.Println(a, b)
			}
		}
	}

}

func DivisoriPropri(n int) (divisori []int) {
	for i := 1; i < n; i++ {
		if n%i == 0 {
			divisori = append(divisori, i)
		}
	}
	return
}

func ÈPerfetto(n int) bool {
	sommaDivisori := 0
	for _, divisore := range DivisoriPropri(n) {
		sommaDivisori += divisore
	}
	return sommaDivisori == n
}

func DivisoriComuni(divisori1, divisori2 []int) (divisoriComuni int) {
	for _, divisore1 := range divisori1 {
		for _, divisore2 := range divisori2 {
			if divisore1 == divisore2 {
				divisoriComuni++
			}
		}
	}
	return
}
