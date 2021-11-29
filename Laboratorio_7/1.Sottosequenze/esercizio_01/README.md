# Esercizio 1

Scrivere un programma che legga da **riga di comando** una stringa composta di caratteri unicode e stampi a schermo tutte le sottostringhe palindrome (che siano uguali lette da destra e da sinistra) della stringa.

##### Esempio d'esecuzione
```text
$ go run esercizio_1.go sottotono
[otto tt tot oto ono]

$ go run esercizio_1.go banana
[ana anana nan ana]

$ go run esercizio_1.go ereggere
[ere ereggere regger egge gg ere]

$ go run esercizio_1.go Vaðlaheiðarvegavinnuverkfærageymsluskúrslyklakippuhringurinn
[nn pp nn]

$ go run esercizio_1.go donaudampfschifffahrtselektrizitätenhauptbetriebswerkbauunterbeamtengesellschaft
[ff fff ff ele izi tät uu ese ll]
```
