# Divisori

Scrivere un programma che legga da **standard input** un numero intero `n` e stampi a video i **divisori propri** del numero letto, ovvero tutti i suoi divisori escluso il numero stesso. 
Ad esempio, i **divisori** del numero 12 sono: 1, 2, 3, 4, 6, 12;
quindi i **divisori propri** di 12 sono: 1, 2, 3, 4, 6.

*Suggerimento:* per controllare se un numero `n` è divisibile per un numero `div`, potete utilizzare l'operatore `%` (`n % div` calcola il della divisione in cui `num` è il dividendo e `div` è il divisore); se `n % div == 0`, allora `div` è un divisore di `n`. 

##### Esempio d'esecuzione:

```text
$ go run divisori.go 
Inserisci numero: 6
Divisori di 6: 1 2 3 

$ go run divisori.go
Inserisci numero: 10
Divisori di 10: 1 2 5 
```