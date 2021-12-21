package main

import (
	trg "example.com/main/triangolo"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	n, _ := strconv.Atoi(os.Args[1])

	triangoli := GeneraTriangoli(n)
	areaMaggiore := AreaMaggiore(triangoli)
	perimetroMinore := PerimetroMinore(triangoli)

	fmt.Println("Triangolo con area maggiore =", trg.String(areaMaggiore))
	fmt.Println("Triangolo con perimetro minore =", trg.String(perimetroMinore))
}

func GeneraTriangoli(n int) (triangoli []trg.Triangolo) {
	rand.Seed(int64(time.Now().Nanosecond()))

	for i := 0; i < n; i++ {
		l1 := rand.Float64()*990 + 10
		l2 := rand.Float64()*990 + 10
		l3 := rand.Float64()*990 + 10
		if t, ok := trg.NuovoTriangolo(l1, l2, l3); ok {
			triangoli = append(triangoli, t)
		}
	}
	return
}

func AreaMaggiore(triangoli []trg.Triangolo) (maggiore trg.Triangolo) {
	for _, t := range triangoli {
		if trg.Area(t) > trg.Area(maggiore) {
			maggiore = t
		}
	}
	return
}

func PerimetroMinore(triangoli []trg.Triangolo) (minore trg.Triangolo) {
	if len(triangoli) > 0 {
		minore = triangoli[0]
		for _, t := range triangoli[1:] {
			if trg.Perimetro(t) < trg.Perimetro(minore) {
				minore = t
			}
		}
	}
	return
}
