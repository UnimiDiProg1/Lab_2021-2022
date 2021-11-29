package main

import (
	"fmt"
	"os"
)

func Palindroma(b []rune) bool {
	for i:=0;i<len(b)/2;i++ {
		if b[i] != b[len(b)-i-1] {
			return false
		}
	}
	
	return true
}

func main() {
	runes := []rune(os.Args[1])

	output := []string{}
	
	for i:=0;i<len(runes)-1;i++ {
		for j:=i+1;j<len(runes);j++ {
			if Palindroma(runes[i:j+1]) {
				output=append(output,string(runes[i:j+1]))
			}
		}
	}

	fmt.Println(output)
}