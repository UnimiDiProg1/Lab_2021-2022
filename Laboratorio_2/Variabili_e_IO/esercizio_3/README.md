# Qual è l'output?

Supponendo che l'input fornito da **standard input** sia 
```
5 6
```
cosa dovrebbe produrre in output il seguente programma?

```go
package main

import "fmt"

func main() {

	var a, b int

	fmt.Scan(&a, &b)
    
    var r int
    
	r = a - b

	fmt.Println(r)

}
```
