# Mazzo di carte

Scrivere un programma che:
1. Legga da **riga di comando** un numero intero `n` tale che `0 < n < 10`.
2. Inizializzi una stringa che rappresenti un mazzo di carte formato dalle sole carte di cuori.
i) Le carte di cuori corrispondono ai caratteri con codice Unicode compreso nell'intervallo che va da `'\U0001F0B1'` a `'\U0001F0BA'`, estremi inclusi. 
ii) Le carte del mazzo non sono mischiate: la prima carta del mazzo è l'asse di cuori; la seconda carta del mazzo è il due di cuori;... l'ultima carta del mazzo è il dieci di cuori.
3. Simuli l'estrazione casuale (ed in sequenza) di `n` carte dal mazzo, stampando a video le carte estratte e quelle rimaste nel mazzo dopo ogni estrazione. 

Oltre alla funzione `main()`, devono essere definite ed utilizzate almeno le seguenti funzioni:

* una funzione `LeggiNumero() int` che legge da **riga di comando** un numero intero e ne restituisce il valore;
* una funzione `GeneraMazzo() string` che restituisce un valore `string` che rappresenta un mazzo di carte formato dalle sole carte di cuori;
* una funzione `EstraiCarta(mazzo string) (cartaEstratta rune, mazzoResiduo string)` che riceve in input un valore `string` nel parametro `mazzo` e restituisce un valore di tipo `rune` nella variabile `cartaEstratta` ed un valore di tipo `string` nella variabile `mazzoResiduo`, che rappresentano rispettivamente la carta estratta in modo casuale dal mazzo di carte rappresentato da `mazzo` ed il mazzo residuo di carte dopo l'estrazione;
* una funzione `EstraiCarte(mazzo string, n int)` che riceve in input un valore `string` nel parametro `mazzo` ed un valore `int` nel parametro `n` e simula l'estrazione casuale (ed in sequenza) di `n` carte dal mazzo di carte rappresentato da `mazzo`, stampando a video le carte estratte e quelle rimaste nel mazzo dopo ogni estrazione; la funzione deve utilizzare la funzione `EstraiCarta()`.

*Suggerimento:* per generare in modo casuale un numero intero, potete utilizzare le funzioni dei package `math/rand` e `time` come mostrato nel seguente frammento di codice:
```go
/* inizializzazione del generatore di numeri casuali */
rand.Seed(int64(time.Now().Nanosecond())) 
/* generazione di un numero casuale compreso nell'intervallo 
   che va da 0 a 99 (estremi inclusi) */
numeroGenerato := rand.Intn(100)
```

##### Esempio d'esecuzione:

```text
$ go run carte.go 6
Estratta la carta 🂱 - Carte rimaste nel mazzo: 🂲🂳🂴🂵🂶🂷🂸🂹🂺
Estratta la carta 🂸 - Carte rimaste nel mazzo: 🂲🂳🂴🂵🂶🂷🂹🂺
Estratta la carta 🂹 - Carte rimaste nel mazzo: 🂲🂳🂴🂵🂶🂷🂺
Estratta la carta 🂷 - Carte rimaste nel mazzo: 🂲🂳🂴🂵🂶🂺
Estratta la carta 🂵 - Carte rimaste nel mazzo: 🂲🂳🂴🂶🂺
Estratta la carta 🂴 - Carte rimaste nel mazzo: 🂲🂳🂶🂺
```
