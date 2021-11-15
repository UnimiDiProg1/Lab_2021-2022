# Retta

Scrivere un programma che legga da **standard input** un valore intero e due valori reali:
- il primo valore è il seme (seed) `s` da utilizzare per inizializzare il generatore di numeri casuali;
- il secondo ed il terzo valore sono il coefficiente angolare `m` e il termine noto `q` di una retta `r: y = m*x + q` sul piano cartesiano.

Una volta terminata la fase di lettura, il programma deve generare per 10 volte una coppia di valori reali `px` e `py` che rappresentano rispettivamente l'ascissa e l'ordinata di un punto sul piano cartesiano e, per ciascun punto:
 1. determinare se, a meno di una costante `EPSILON = 1.0e-9`, il punto sta sopra, sotto, o appartiene alla retta `r`;
 2. stampare a video il relativo messaggio (come mostrato nell'**Esempio di esecuzione**).

I valori `px` e `py` devono essere compresi nell'intervallo [0, 10.0).

Oltre alla funzione `main()`, il programma deve utilizzare almeno due tra le seguenti funzioni:

```go
const EPSILON = 1.0e-9

/* La funzione `func ÈXMaggioreDiY(x, y float64) bool` riceve in input due 
valore reali nei parametri `x` e `y` e restituisce `true` se `x` è maggiore 
di `y` di almeno una costante EPSILON */
/* ÈXMaggioreDiY(5.0,4.999) -> true */
/* ÈXMaggioreDiY(5.0,4.9999999999) -> false */
func ÈXMaggioreDiY(x, y float64) bool {
    return (x - y) > EPSILON 
}

/* La funzione `func ÈXUgualeAY(x, y float64) bool` riceve in input due 
valore reali nei parametri `x` e `y` e restituisce `true` se `x` è uguale 
a `y` a meno di una costante EPSILON */
/* ÈXUgualeAY(5.0,4.999) -> false */
/* ÈXUgualeAY(5.0,4.9999999999) -> true */
func ÈXUgualeAY(x, y float64) bool {
    return math.Abs(x - y) <= EPSILON
}

/* La funzione `func ÈXMaggioreDiY(x, y float64) bool` riceve in input due 
valore reali nei parametri `x` e `y` e restituisce `true` se `x` è minore 
di `y` di almeno una costante EPSILON */
/* ÈXMinoreDiY(4.999,5.0) -> true */
/* ÈXMinoreDiY(4.9999999999,5.0) -> false */
func ÈXMinoreDiY(x, y float64) bool {
    return (x - y) < -EPSILON
}
```

*Nota:* Inizializzando il generatore dei numeri casuali con l'istruzione
```go
/* inizializzazione del generatore di numeri casuali */
rand.Seed(s) 
```
la sequenza dei numeri casuali generati durante l'esecuzione del programma dipenderà dal valore `s`, quindi due esecuzioni successive del programma considereranno la stessa sequenza di numeri casuali se l'utente inserirà lo stesso valore per `s`.

*Suggerimento:* per generare in modo casuale un valore reale, potete utilizzare la funzione `func Float64()` del package `math/rand`, che restituisce un valore casuale compreso tra [0.0,1.0).

##### Esempio d'esecuzione:

```markdown
$ go run retta.go
Inserisci s: 2
Inserisci m e q: 1 0

Punto (0.8364831721292811,1.325271527168901) - Il punto sta sopra la retta
Punto (0.25744012651422055,0.5948870460174552) - Il punto sta sopra la retta
Punto (3.071353082885974,4.638274793965569) - Il punto sta sopra la retta
Punto (2.125002747927654,1.030771445518701) - Il punto sta sotto la retta
Punto (1.059760133310769,0.6344848690006494) - Il punto sta sotto la retta
Punto (3.09889044262683,4.311133779707338) - Il punto sta sopra la retta
Punto (1.843213441859697,1.9013752063486216) - Il punto sta sopra la retta
Punto (3.1086212867763456,4.552073840109743) - Il punto sta sopra la retta
Punto (2.151010640152745,0.06284952382418317) - Il punto sta sotto la retta
Punto (2.6797238953528506,4.4750914842551826) - Il punto sta sopra la retta

$ go run retta.go
Inserisci s: 2
Inserisci m e q: 1 1

Punto (0.8364831721292811,1.325271527168901) - Il punto sta sotto la retta
Punto (0.25744012651422055,0.5948870460174552) - Il punto sta sotto la retta
Punto (3.071353082885974,4.638274793965569) - Il punto sta sopra la retta
Punto (2.125002747927654,1.030771445518701) - Il punto sta sotto la retta
Punto (1.059760133310769,0.6344848690006494) - Il punto sta sotto la retta
Punto (3.09889044262683,4.311133779707338) - Il punto sta sopra la retta
Punto (1.843213441859697,1.9013752063486216) - Il punto sta sotto la retta
Punto (3.1086212867763456,4.552073840109743) - Il punto sta sopra la retta
Punto (2.151010640152745,0.06284952382418317) - Il punto sta sotto la retta
Punto (2.6797238953528506,4.4750914842551826) - Il punto sta sopra la retta
``` 