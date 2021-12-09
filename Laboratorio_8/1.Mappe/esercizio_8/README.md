# Sottosequenze (2)

Scrivere un programma che:
* legga da **riga di comando** una sequenza `s` di valori che rappresentano caratteri appartenenti all'alfabeto inglese (e quindi codificati all'interno dello standard US-ASCII (integrato nello standard Unicode));
* stampi a video tutte le sottosequenze di caratteri presenti in `s` che:
  1. iniziano e finiscono con lo stesso carattere; 
  2. sono formate da almeno 3 caratteri.

Ciascuna sottosequenza deve essere stampata un'unica volta, riportando il relativo numero di occorrenze della sottosequenza in `s` (cfr. **Esecuzione d'esecuzione**).

Se non esistono sottosequenze che soddisfano le condizioni 1 e 2, il programma non deve stampare nulla.

Si noti che una sottosequenza può essere contenuta in un'altra sottosequenza più grande.

Si assuma che la sequenza di valori specificata a riga di comando sia nel formato corretto e includa almeno 3 caratteri.

##### Esempio d'esecuzione:

```text
$ go run sottosequenze.go a b b a b b a
a b b a -> Occorrenze: 2
a b b a b b a -> Occorrenze: 1
b b a b -> Occorrenze: 1
b b a b b -> Occorrenze: 1
b a b -> Occorrenze: 1
b a b b -> Occorrenze: 1

$ go run sottosequenze.go a b c a c b a
a b c a c b a -> Occorrenze: 1
b c a c b -> Occorrenze: 1
c a c -> Occorrenze: 1
a c b a -> Occorrenze: 1
a b c a -> Occorrenze: 1

$ go run sottosequenze.go e a b c a c f
c a c -> Occorrenze: 1
a b c a -> Occorrenze: 1

$ go run sottosequenze.go e a b b c a b c b f
a b b c a -> Occorrenze: 1
b b c a b -> Occorrenze: 1
b b c a b c b -> Occorrenze: 1
b c a b -> Occorrenze: 1
b c a b c b -> Occorrenze: 1
c a b c -> Occorrenze: 1
b c b -> Occorrenze: 1

$ go run sottosequenze.go a b c c e

```

Siccome le mappe non sono strutture ordinate, il vostro output potrà presentare le sottosequenze in ordine diverso. L'importante è che ci siano tutte e con il giusto numero di occorrenze.