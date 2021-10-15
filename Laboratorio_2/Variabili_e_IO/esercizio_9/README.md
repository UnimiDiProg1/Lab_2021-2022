# Età

Scrivere un programma che legga da **standard input** le età di due persone (espresse in anni) e calcoli:

* la somma delle età inserite;
* la media delle età inserite;
* la media delle età inserite arrotondata per difetto all'intero inferiore;
* la media delle età inserite arrotondata per eccesso all'intero superiore;
* la somma e la media delle età che le due persone avranno fra 10 anni.

*Suggerimento:* 
la media arrotondata per difetto può essere calcolata usando la funzione `math.Floor` del package `math` nel seguente modo:
```go
var mediaArrotondataDifetto float64 = math.Floor(media)
```
Similarmente, la media arrotondata per eccesso può essere calcolata usando la funzione `math.Ceil` nel seguente modo:
```go
var mediaArrotondataEccesso float64 = math.Ceil(media)
```

##### Esempio d'esecuzione:

```text
$ go run calcoloeta.go 
Età persona 1 = 15
Età persona 2 = 20
Somma età = 35
Media età = 17.5
Media età arrotondata per difetto all'intero inferiore = 17
Media età arrotondata per eccesso all'intero superiore = 18
Somma età a 10 anni = 55
Media età a 10 anni = 27.5
```
