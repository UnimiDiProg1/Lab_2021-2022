# Triangoli casuali

Si estenda il package `triangolo` definito relativamente all'esercizio **1 Perimetro e Area di un triangolo** implementando la funzione:

* `String(t Triangolo) string` che riceve in input un'instanza del tipo `Triangolo` nel parametro `t` e restituisce un valore `string` che corrisponde alla rappresentazione `string` di `t` nel formato `Triangolo con lati L1, L2 e L3.`, dove `L1`, `L2` ed `L3` sono i valori ai campi `lato1`, `lato2` e `lato3` di `t`. La conversione in `string` dei valori dei tre lati deve essere effettuata utilizzando due cifre decimali.

Utilizzando le funzionalità messe a disposizione dal package `triangolo`, scrivere un programma che:
* legga da **riga di comando** un numero intero `n`;
* generi in maniera casuale `n` triple di valori reali compresi tra `10` e `1000`; i valori `l1`, `l2`, `l3` di ciascuna tripla corrispondono alle misure dei lati di un ipotetico triangolo;
* stampi a video la rappresentazione `string` del triangolo con area più grande tra quelli corrispondenti alle triple di valori reali generate; 
* stampi a video la rappresentazione `string` del triangolo con perimetro più piccolo tra quelli corrispondenti alle triple di valori reali generate.

Oltre alla funzione `main()`, devono essere definite ed utilizzate almeno le seguenti funzioni:

* una funzione `GeneraTriangoli(n int) (tN []triangolo.Triangolo)` che riceve in input un valore `int` nel parametro `n` e restituisce un valore `[]triangolo.Triangolo` nella variabile `tN` in cui sono memorizzate solamente le istanze di tipo `triangolo.Triangolo` *valide*.

##### Esempio d'esecuzione:

```text
$ go run triangoli_casuali.go 10
Triangolo con area maggiore = Triangolo con lati 716.05, 867.05 e 976.47
Triangolo con perimetro minore = Triangolo con lati 590.89, 266.73 e 559.86
```
