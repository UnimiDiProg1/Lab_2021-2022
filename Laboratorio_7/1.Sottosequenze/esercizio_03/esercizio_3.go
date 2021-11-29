package main

import (
	"fmt"
	"os"
)

func Sottostringhe(s string) []string {
	result := []string{}

	for i:=0; i<len(s)-1;i++ {
		for j:=i+1; j<len(s);j++ {
			if s[j] <= s[j-1] {
				break
			}

			if j != i {
				result = append(result,s[i:j+1])
			}
		}
	}
	
	return result
}

func main() {
	
	for _, c := range os.Args[1] {
		if c < '0' || c > '9' {
			return
		}
	}
	
	sottostringhe := Sottostringhe(os.Args[1])
	
	fmt.Println(sottostringhe)
}