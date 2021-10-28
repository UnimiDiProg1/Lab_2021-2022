# Quadrato con diagonale

Scrivere un programma che legga da **standard input** un numero intero `n` e che, come mostrato nell'**Esempio di esecuzione**, stampi a video un quadrato di lato `n` in cui:
* una diagonale è rappresentata utilizzando il carattere `o` (lettera o);
* l'area del quadrato al di sotto della diagonale è rappresentata con il carattere `*` (asterisco);
* l'area del quadrato al di sopra della diagonale è rappresentata con il carattere `+` (più);
* i caratteri di ogni riga del quadrato sono separati tra di loro da un carattere spazio.

##### Esempio d'esecuzione:

```text
$ go run quadrato_con_diagonale.go
Inserisci un numero: 7
o + + + + + + 
* o + + + + + 
* * o + + + + 
* * * o + + + 
* * * * o + + 
* * * * * o + 
* * * * * * o 
```