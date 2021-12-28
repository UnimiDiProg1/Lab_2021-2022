package main

import (
	"fmt"
	"sort"
)

func Sottostringhe(s string) map[string]int {
	result := make(map[string]int)

	for i:=0; i<len(s)-1;i++ {
		for j:=i+1; j<len(s);j++ {
			if s[j]>s[j-1]{
				substring := s[i:j+1]
				result[substring]++
			} else {
				break
			}
		}
	}

	return result
}

func main() {
	var s string
	fmt.Scanln(&s)
	
	for _, c := range s {
		if c < 'a' || c > 'z' {
			return
		}
	}
	
	sottostringhe := Sottostringhe(s)
	
	//BONUS PART
	var keys []string 

	for k,_ := range sottostringhe {
		keys = append(keys,k)
	}

	sort.Strings(keys)

	fmt.Println("output:")
	
	for i := range keys{
		key := keys[i]
		fmt.Println(key, sottostringhe[key])
	}
}