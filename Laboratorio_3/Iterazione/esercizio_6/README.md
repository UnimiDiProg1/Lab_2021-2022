# Operazioni con un numero variabile di valori

Scrivere un programma che, dopo aver letto da **standard input** un numero intero `n`, chiede all'utente di inserire `n` numeri interi (sempre da **standard input**). 

Dopo aver letto gli `n` numeri interi, il programma deve stampare:
* la somma degli `n` numeri letti;
* il minimo tra i numeri letti;
* il massimo tra i numeri letti;
* il numero di interi letti strettamente positivi (maggiori di 0), strettamente negativi (minori di 0), e nulli.

*Suggerimento*: per leggere `n` numeri, potete utilizzare un ciclo `for` che per `n` volte utilizza `fmt.Scan`. 

##### Esempio d'esecuzione:

```text
$ go run nnumeri.go
9
Inserisci 9 numeri:
1 -2 3 -4 5 -6 7 -8 9
somma = 5
valore minimo = -8
valore massimo = 9
interi > 0 = 5
interi < 0 = 4
interi = 0 = 0
```