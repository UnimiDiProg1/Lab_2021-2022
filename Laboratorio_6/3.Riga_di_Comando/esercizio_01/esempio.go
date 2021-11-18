package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("TIPO os.Args: %T\n", os.Args)
	for i, v := range os.Args {
		fmt.Printf("os.Args[%d]: TIPO = %T - VALORE = %s\n", i, v, v)
	}
}
