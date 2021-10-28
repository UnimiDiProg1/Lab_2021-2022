# Radice quadrata

Scrivere un programma che legga da **standard input** un numero reale `n` e stampi a video il valore della radice quadrata di `n` solo se `n >= 0`. Se `n < 0` il programma deve stampare `Il valore inserito non appartiene al dominio della funzione`.

*Suggerimento:* Per calcolare la radice quadrata di un numero relae `n` potete usare la funzione `Sqrt` del package `math`:
```go
radiceQuadrata = math.Sqrt(n)
```

Oltre alla funzione `main()`, il programma deve definire e utilizzare le seguenti funzioni:
* `RadiceQuadrata(numero float64) (float64, bool)` che riceve in input un valore reale nel parametro `numero`. Se `numero >= 0` la funzione restituisce il valore della radice quadrata di `numero` e un valore logico `true` come certificato della corretta esecuzione del calcolo. Altrimenti, la funzione restituisce un valore reale `0` e un valore logico `false`, per segnalare che non è stato possibile calcolare la radice quadrata di `numero` nel campo dei reali.

##### Esempio d'esecuzione:

```text
$ go run radice.go 
Inserisci un numero: 10
Radice quadrata: 3.1622776601683795

$ go run radice.go
Inserisci un numero: 4
Radice quadrata: 2

$ go run radice.go
Inserisci un numero: -1
Il valore inserito non appartiene al dominio della funzione
```