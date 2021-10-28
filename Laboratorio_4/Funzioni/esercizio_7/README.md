# Cerchio

Scrivere un programma che legga da **standard input** un numero reale `raggio` e stampi i valori di circonferenza e area di un cerchio di raggio `raggio`.

Oltre alla funzione `main()`, il programma deve definire e utilizzare le seguenti funzioni:
* `CalcolaArea(raggio float64) float64` che riceve in input il valore reale del raggio di un cerchio nel parametro `raggio` e restituisce il valore dell'area del cerchio associato;
* `CalcolaCirconferenza(raggio float64) float64` che riceve in input il valore reale del raggio di un cerchio nel parametro `raggio` e restituisce il valore della circonferenza del cerchio associato.

##### Esempio d'esecuzione:

```text
$ go run cerchio.go 
Inserisci il raggio del cerchio: 10.5
Area del cerchio: 346.36059005827474
Circonferenza del cerchio: 65.97344572538566

$ go run cerchio.go 
Inserisci il raggio del cerchio: 3.6
Area del cerchio: 40.71504079052372
Circonferenza del cerchio: 22.61946710584651
```
