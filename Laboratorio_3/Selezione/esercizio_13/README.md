# Retta

Scrivere un programma che legga da **standard input** 4 valori a virgola mobile:  
- i primi due valori sono il coefficiente angolare `m` e il termine noto `q` di una retta `r: y = m*x + q`
- il terzo e il quarto valore sono le coordinate `px` e `py` di un punto `P(px,py)`

Il programma deve determinare se il punto `P` sta sopra o sotto la retta od appartiene ad essa, e stampare a video il relativo messaggio.

*Suggerimento:* un punto appartiene ad una retta se sostituendo le sue coordinate nell'equazione della retta l'uguaglianza è verificata. Un punto sta sopra una retta se sostituendo il valore dell'ascissa nell'equazione della retta si ottiene `y < py`.

```text
$ go run retta.go
Inserisci m e q: 1 0
Inserisci x e y: 5 5
Il punto appartiene alla retta

$ go run retta.go
Inserisci m e q: 1 1
Inserisci x e y: 5 5
Il punto sta sotto la retta
```