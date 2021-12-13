# Perimetro e Area di un triangolo
Scrivere un programma che:
* legga da **riga di comando** tre valori reali che corrispondono alle misure dei lati di un triangolo;
* stampi a video il valore del perimetro e dell'area del triangolo corrispondente ai tre valori reali letti.

Il programma deve utilizzare le funzionalità messe a disposizione da un package locale `triangolo` in cui è definito il tipo `Triangolo`:
```go
type Triangolo struct {
    lato1, lato2, lato3 float64
}
```
e le seguenti funzioni:
* una funzione `NuovoTriangolo(l1, l2, l3 float64) (t Triangolo, ok bool)` che, se `l1+l2 > l3`, `l1+l3 > l2` e `l2+l3 > l1`, restituisce una nuova istanza del tipo `Triangolo` inizializzata in base ai valori dei parametri `l1`, `l1` e `l3` (nella variabile `t`) ed il valore `true` (nella variabile `ok`); altrimenti un Triangolo vuoto e il valore `false` nella variabile ok;
* una funzione `Perimetro(t Triangolo) float64` che riceve in input un'instanza del tipo `Triangolo` nel parametro `t` e restituisce un valore `float64` pari al perimetro del triangolo rappresentato da `t`;
* una funzione `Area(t Triangolo) float64` che riceve in input un'instanza del tipo `Triangolo` nel parametro `t` e restituisce un valore `float64` pari all'area del triangolo rappresentato da `t`, calcolato utilizzando la formula di Erone:
```go
p := (lato1 + lato2 + lato3) / 2
area := math.Sqrt(p * (p-lato1) * (p-lato2) * (p-lato3))
```

## Creazione e utilizzo di moduli
Un package creato localmente è posizionato in una nuova cartella all'interno della cartella principale del progetto. Il package e la sua cartella devono avere lo stesso nome. Ad esempio il package `triangolo` verrà messo nella cartella `triangolo`.

Al fine di importare il nuovo package creato, bisogna inizializzare il modulo corrispondente al progetto sul quale si sta lavorando. Questo si fa richiamando il comando `go mod init <url del modulo>` dove l'url del modulo è la repository del modulo (se esiste una), altrimenti si può usare un identificatore fittizio, ad esempio `example.com/main`.

Per importare il package `triangolo` a questo punto bisognerà scrivere nel package main `import example.com/main/triangolo`.

##### Esempio di creazione del modulo:
```text
$ go mod init example.com/main
go: creating new go.mod: module example.com/main
go: to add module requirements and sums:
        go mod tidy
```

##### Esempio d'esecuzione:

```text
$ go run area_e_perimetro.go 5 4 3
Perimetro triangolo = 12
Area triangolo = 6

$ go run area_e_perimetro.go 4 4 4
Perimetro triangolo = 12
Area triangolo = 6.928203230275509

$ go run area_e_perimetro.go 10 3 5
Errore: impossibile creare un triangolo con le misure specificate
```
