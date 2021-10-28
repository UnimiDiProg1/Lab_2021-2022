# Carte

Sapendo che al codice Unicode 127153 (associato alla rappresentazione in bit Unicode/UTF-8 `'\U0001F0B1'`) corrisponde il simbolo "asso di cuori", e che i codici successivi corrispondono alle carte successive (2 di cuori, 3 di cuori, ...), scrivere un programma che stampi tutte le carte da gioco dall'asso di cuori al 10 di cuori.

*Suggerimento:* Un carattere è una variabile di tipo `rune`, il cui valore è un intero corrispondente al codice Unicode del carattere. Le istruzioni equivalenti `var c rune = 127153` e `var c rune = '\U0001F0B1'` servono a definire la varibile `c` di tipo `rune` ed inizializzarla al valore "asse di cuori". Per stampare la carta da gioco "asse di cuori" si può utilizzare l'istruzione `fmt.Print(string(c))` o `fmt.Printf("%c",c)`.

```text
$ go run carte.go 
Simbolo: 🂱 - Codice numerico in base 10: 127153
Simbolo: 🂲 - Codice numerico in base 10: 127154
Simbolo: 🂳 - Codice numerico in base 10: 127155
Simbolo: 🂴 - Codice numerico in base 10: 127156
Simbolo: 🂵 - Codice numerico in base 10: 127157
Simbolo: 🂶 - Codice numerico in base 10: 127158
Simbolo: 🂷 - Codice numerico in base 10: 127159
Simbolo: 🂸 - Codice numerico in base 10: 127160
Simbolo: 🂹 - Codice numerico in base 10: 127161
Simbolo: 🂺 - Codice numerico in base 10: 127162
```
