# Divisione per zero

Si consideri il seguente frammento di codice.

```go
var numero float64 = 0
fmt.Println(numero)

fmt.Println(0 / numero)

numero = 1 / numero
fmt.Println(numero)

numero = 1 / numero
fmt.Println(numero)

numero = -(1 / numero)
fmt.Println(numero)

numero = numero - numero
fmt.Println(numero)
```
Quali valori vengono stampati a video?
Quali valori verrebbero stampati a video se la variabile `numero` fosse di tipo `int`?

