# Fizz Buzz

Scrivere un programma che legga da **standard input** un intero `n`.
Il programma deve stampare a video la sequenza di numeri che vanno da 1 a `n` come mostrato nell'**Esempio d'esecuzione**. In particolare:
 * ogni numero divisibile per 3 deve essere rimpiazzato dalla parola `"Fizz"`;
 * ogni numero divisibile per 5 deve essere rimpiazzato dalla parola `"Buzz"`;
 * ogni numero divisibile sia per 3 sia per 5 deve essere sostituito da `"Fizz Buzz"`.

##### Esempio d'esecuzione:

```text
$ go run fizzbuzz.go 
6
1 2 Fizz 4 Buzz Fizz 

$ go run fizzbuzz.go
20
1 2 Fizz 4 Buzz Fizz 7 8 Fizz Buzz 11 Fizz 13 14 FizzBuzz 16 17 Fizz 19 Buzz 
```