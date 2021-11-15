# Overflow

Qual è l'output del seguente programma?

```go
import "fmt"

func main() {
	var (
		a, b int8 = 30, 100
	)
	somma := a + b
	fmt.Println("La somma di", a, "e", b, "è", somma)
}
```
