# Uguaglianza tra valori

Scrivere un programma che:  
1) Legga da **standard input** un numero reale.  
2) Calcoli la radice quadrata del numero letto (sia `x` la variabile di tipo `float64` in cui è memorizzato il valore; la radice quadrata di del valore reale può essere calcolata utilizzando la funzione `math.Sqrt` del package `math` nel seguente modo `radiceQuadrata := math.Sqrt(x)`).  
3) Stampi a video `x, "uguale a", radiceQuadrata, "*", radiceQuadrata, "\n"` nel caso in cui `radiceQuadrata*radiceQuadrata` sia uguale a `x` e `x, "diverso da", radiceQuadrata, "*", radiceQuadrata, "\n"` altrimenti.  

Supponendo che l'utente inserisca da **standard input** il valore `9`,
qual è l'output del programma?

Supponendo che l'utente inserisca da **standard input** il valore `10`,
qual è l'output del programma? Perché?
 
##### Esempio d'esecuzione:

```markdown
$ go run uguaglianza.go
Inserisci un numero: 4
4 uguale a 2 * 2
``` 