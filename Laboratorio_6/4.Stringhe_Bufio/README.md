# Qual è l'output?

Qual è l'output di questo programma?

```go
package main

import "fmt"

func main() {

	var s string = "Ma il cielo è sempre più blu!"

	var s2 []rune = []rune(s)

	fmt.Println(len(s), len(s2))
}
```
# Qual è l'output?

Qual è l'output del seguente programma?

```go
package main

import "fmt"

func main() {

	//strFrom := "ABDEF"
	//strFrom[2] = 'C' /* error: "cannot assign to strFrom[2]" */

	/* (1) */
	strFrom1 := "ABÈEF"
	// len(sliceDiRune) è pari al numero di caratteri Unicode presenti in strFrom1
	sliceDiRune := []rune(strFrom1)
	sliceDiRune[2] = 'C'
	strTo1 := string(sliceDiRune)
	fmt.Println(strTo1)
	/* (1) - end */

	/* (2) */
	strFrom2 := "ABÈEF"
	//strTo2 := strFrom2[:2]+string('C')+strFrom2[3:]
	strTo2 := strFrom2[:2] + "C" + strFrom2[3:]
	fmt.Println(strTo2)
	/* (2) - end */
}
```
# Controlla sequenza (1)

Scrivere un programma che legga da **standard input** una stringa di caratteri in cui i caratteri corrispondenti a cifre decimali rappresentano dei numeri interi minori di 10.

All'interno della stringa letta, i caratteri corrispondenti a cifre decimali sono intervallati tra loro da sottostringhe formate da caratteri arbitrari.

Il programma deve considerare la sequenza di numeri interi nascosta all'interno della stringa letta da **standard input**.

**Esempio:**

All'interno della stringa di caratteri

`&4&$4%mamma5!6mia6cosa1succede0`

la sequenza di numeri interi nascosta è:

`4 4 5 6 6 1 0`

Il programma deve controllare che i numeri presenti all'interno della sequenza siano ordinati in senso decrescente, cioè dal più grande al più piccolo.

Nel caso in cui i numeri interi siano ordinati in senso decrescente, il programma deve stampare a video il messaggio `Sequenza nascosta ordinata.`.
In caso contrario, il programma deve stampare a video il messaggio `Sequenza nascosta non ordinata.`.

Si assuma che:
* all'interno della stringa di caratteri letta da **standard input** non siano presenti caratteri consecutivi corrispondenti a cifre decimali;
* all'interno della stringa di caratteri letta da **standard input** non sia presente nessun carattere di spaziatura, ossia un carattere il cui codice Unicode, passato come argomento alla funzione `func IsSpace(r rune) bool` del package `unicode`, fa restituire `true` alla funzione.

##### Esempio d'esecuzione:
```text
$ go run controlla_sequenza.go
abc8dèf6asa2sd0
Sequenza nascosta ordinata.

$ go run controlla_sequenza.go
Èbc8d8e5a#4§d1e
Sequenza nascosta ordinata.

$ go run controlla_sequenza.go
a°c9d8e6aà7sd1e
Sequenza nascosta non ordinata.
```
# Testo concatenato

Scrivere un programma che: 
* legga da **standard input** un testo su più righe;
* termini la lettura quando, premendo la combinazione di tasti `Ctrl+D`, viene inserito l'indicatore End-Of-File (EOF);
* ristampi il testo letto.

**Nota:** la stampa del testo deve essere effettuata solamente dopo aver letto per intero il testo.  

Oltre alla funzione `main()`, il programma deve definire e utilizzare le seguenti funzioni:
* una funzione `LeggiTesto() string` che legge da **standard input** un testo su più righe e terminato dall'indicatore EOF, restituendo un valore `string` in cui è memorizzato il testo letto.

##### Esempio d'esecuzione:

```text
$ go run ristampa.go 
Inserisci testo (termina con CTRL+D):
Testo di prova
disposto su più righe

Alcune possono essere anche righe vuote
Testo letto:
Testo di prova
disposto su più righe

Alcune possono essere anche righe vuote
```
# Statistiche testo

Scrivere un programma che: 
* legga da **standard input** un testo su più righe;
* termini la lettura quando, premendo la combinazione di tasti `Ctrl+D`, viene inserito da **standard input** l'indicatore End-Of-File (EOF);
* stampi a video le seguenti statistiche relative al testo letto:
  1. il numero di parole presenti nel testo;
  2. la lunghezza media delle parole presenti nel testo.

**Nota:** una parola è una sequenza di caratteri consecutivi rappresentanti delle lettere. Numeri, punteggiatura e caratteri di spaziatura intervallano parole diverse.

Oltre alla funzione `main()`, il programma deve definire e utilizzare le seguenti funzioni:
* una funzione `LeggiTesto() string` che legge da **standard input** un testo su più righe terminato dall'indicatore EOF (`CTRL+D`), restituendo un valore `string` in cui è memorizzato il testo letto;
* una funzione `StatisticheParole(s string) (int, int)` che riceve in input un valore `string` nel parametro `s` e restituisce due valori `int` pari rispettivamente al numero di parole presenti in `s` e alla loro lunghezza totale.

*Suggerimento:* per sapere se un carattere rappresenta una lettera utilizza la funzione `unicode.IsLetter()` del package `unicode`. Utilizza `go doc unicode.IsLetter` per scoprire il suo funzionamento.

##### Esempio d'esecuzione:

```text
$ go run statistiche.go    
Inserisci un testo su più righe (termina con Ctrl+D):
Testo di prova                      

su cui calcolare le statistiche.
Test 01: prova.prova0prova
Statistiche:
Numero parole: 12
Lunghezza media: 4.833333333333333
```
# Testo invertito

Scrivere un programma che legga da **standard input** un testo formato da un numero variabile di righe, terminando la lettura quando viene inserita da **standard input** una riga vuota (`""`), e lo ristampi dall’ultimo carattere al primo.

Scrivere un programma che: 
* legga da **standard input** un testo su più righe;
* termini la lettura quando viene inserita da **standard input** una riga vuota (`""`);
* ristampi il testo letto (riga vuota esclusa) dall’ultimo carattere al primo.

Oltre alla funzione `main()`, il programma deve definire e utilizzare le seguenti funzioni:
* una funzione `LeggiTesto() string` che legge da **standard input** un testo su più righe e terminato da una riga vuota (`""`), restituendo un valore `string` in cui è memorizzato il testo letto (riga vuota esclusa);
* una funzione `InvertiStringa(s string) string` che riceve in input un valore `string` nel parametro `s` e ne inverte l'ordine dei caratteri, ovvero restituisce un valore `string` in cui il primo carattere è l'ultimo che definisce `s`, il secondo carattere è il penultimo che definisce `s`, ...  e l'ultimo carattere è il primo che definisce `s`.

##### Esempio d'esecuzione:

```text
$ go run stringainvertita.go
Inserisci un testo su più righe (termina con riga vuota):
Testo di prova
disposto su due righe

Testo invertito:
ehgir eud us otsopsid
avorp id otseT
```
# Numero nascosto

Scrivere un programma che legga da **standard input** una riga di testo e stampi in output il doppio del numero nascosto all'interno della riga di testo, ovvero il doppio del numero che si ottiene concatenando le cifre presenti all'interno della riga di testo. Il programma non stampa nulla se non è presente alcun numero nascosto.

*Suggerimento:* Per convertire una stringa in un numero intero utilizzate la funzione `strconv.Atoi()` del package `strconv`. Invece, per sapere se un carattere è una cifra utilizzate la funzione `unicode.IsDigit()` del package `unicode`.

Oltre alla funzione `main()`, il programma deve definire e utilizzare le seguenti funzioni:
* una funzione `LeggiTesto() string` che legge da **standard input** una riga di testo, restituendo un valore `string` in cui è memorizzato il testo letto;
* una funzione `NumeroNascosto(testo string) (int, error)` che riceve in input un valore `string` nel parametro `testo` e restituisce due valori:
   * il primo valore è un numero intero che rappresenta il numero nascosto all'interno del testo. Se il testo in input non contiene alcun numero il valore restituito deve essere `0`;
   * il secondo valore è l'eventuale errore restituito dalla funzione `strconv.Atoi()`.

##### Esempio d'esecuzione:

```text
$ go run numero_nascosto.go
Ci8ao 97com3 va?
Doppio del numero nascosto: 17946 (8973 * 2)

$ go run numero_nascosto.go
Ch3 831 t3mp0
Doppio del numero nascosto: 766260 (383130 * 2)

$ go run numero_nascosto.go
c140n3
Doppio del numero nascosto: 2806 (1403 * 2)

$ go run numero_nascosto.go
nessun numero nascosto
```
# Linguaggio farfallino

Nel linguaggio farfallino ciascuna vocale non accentata (`vocale`) viene sostituita da una sequenza di tre caratteri `vocale-f-vocale`. Per esempio, la vocale `a` viene sostituita dalla sequenza `afa`, la vocale `e` dalla sequenza `efe` e così via. Se una vocale è maiuscola, anche la sequenza di tre caratteri che sostituisce la vocale deve essere definita da caratteri maiuscoli (ad esempio, la vocale `A` viene sostituita dalla sequenza `AFA`).

Scrivere un programma che: 
* legga da **standard input** un testo su più righe;
* termini la lettura quando, premendo la combinazione di tasti `Ctrl+D`, viene inserito da **standard input** l'indicatore End-Of-File (EOF);
* ristampi il testo letto dopo averlo tradotto in linguaggio farfallino.

Oltre alla funzione `main()`, il programma deve definire e utilizzare le seguenti funzioni:
* una funzione `LeggiTesto() string` che legge da **standard input** un testo su più righe e terminato dall'indicatore EOF, restituendo un valore `string` in cui è memorizzato il testo letto;
* una funzione `TraduciCarattereInFarfallino(c rune) string` che riceve in input un valore `rune` nel parametro `c` e restituisce un valore `string` che rappresenta la traduzione in linguaggio farfallino di `c`;
* una funzione `TraduciTestoInFarfallino(s string) string` che riceve in input un valore `string` nel parametro `s` e restituisce un valore `string` che rappresenta la traduzione in linguaggio farfallino di `s`.

*Suggerimento:* per verificare se un carattere è una vocale potete utilizzare i costrutti `if` o `switch`, oppure utilizzare la funzione `strings.ContainsRune()` del package `strings`.

##### Esempio d'esecuzione:

```text
$ go run farfallino.go 
Inserisci un testo su più righe (termina con CTRL+D):
Questo è un testo di prova
da trasformare IN ALFABETO FARFALLINO
Risultato:
Qufuefestofo èfè ufun tefestofo difi profovafa
dafa trafasfoformafarefe IFIN AFALFAFABEFETOFO FAFARFAFALLIFINOFO
```
# Garibaldi

Scrivere un programma che: 
* legga da **standard input** un testo su più righe;
* termini la lettura quando viene inserita da **standard input** una riga vuota (`""`);
* ristampi il testo letto (riga vuota esclusa) sostituendo tutte le vocali con delle `u`.

Oltre alla funzione `main()`, il programma deve definire e utilizzare le seguenti funzioni:
* una funzione `LeggiTesto() string` che legge da **standard input** un testo su più righe e terminato da una riga vuota (`""`), restituendo un valore `string` in cui è memorizzato il testo letto (riga vuota esclusa);
* una funzione `TrasformaCarattere(c rune) rune` che riceve in input un valore `rune` nel parametro `c` e restituisce un valore `rune` uguale a `u` se `c` è una vocale, e uguale a `c` altrimenti.
* una funzione `Garibaldi(s string) string` che riceve in input un valore `string` nel parametro `s` e restituisce un valore `string` in cui ogni carattere è trasformato utilizzando la funzione `TrasformaCarattere()`.

##### Esempio d'esecuzione:

```text
$ go run garibaldi.go
Inserisci un testo su più righe (termina con riga vuota):
Garibaldi fu ferito
fu ferito in una gamba,
Garibaldi che comanda
che comanda i bersaglier

Risultato trasformazione:
Gurubuldu fu furutu
fu furutu un unu gumbu,
Gurubuldu chu cumundu
chu cumundu u bursugluur
```
# Spaziatura

Scrivere un programma che: 
* legga da **standard input** un testo su più righe;
* termini la lettura quando, premendo la combinazione di tasti `Ctrl+D`, viene inserito da **standard input** l'indicatore End-Of-File (EOF);
* ristampi il testo letto con spaziatura, ovvero inserendo uno spazio `' '` tra ogni coppia di caratteri che **non** sono spazi.

Oltre alla funzione `main()`, il programma deve definire e utilizzare le seguenti funzioni:
* una funzione `LeggiTesto() string` che legge da **standard input** un testo su più righe e terminato dall'indicatore EOF, restituendo un valore `string` in cui è memorizzato il testo letto;
* una funzione `Spazia(s string) string` che riceve in input un valore `string` nel parametro `s` e restituisce un valore `string` che rappresenta la versione spaziata di `s`.

*Suggerimento:* per sapere se un carattere è uno spazio utilizzate la funzione `unicode.IsSpace` del package `unicode`.

##### Esempio d'esecuzione:

```text
$ go run spaziatura.go
Inserisci un testo su più righe (termina con CTRL+D):
Testo di prova 
da stampare con spaziatura
Risultato:
T e s t o d i p r o v a 
d a s t a m p a r e c o n s p a z i a t u r a
```
# Il cifrario di Cesare

Giulio Cesare usava per le sue corrispondenze riservate un codice di sostituzione molto semplice, nel quale la lettera chiara viene sostituita dalla lettera che la segue di tre posti nell’alfabeto: la lettera A è sostituita dalla D, la B dalla E, e così via fino alle ultime lettere che sono cifrate con le prime.

Chiaro:   A B C D E F G H I J K L M N O P Q R S T U V W X Y Z

Cifrato:  D E F G H I J K L M N O P Q R S T U V W X Y Z A B C

Più in generale si dice cifrario di Cesare un codice nel quale ogni lettera del messaggio chiaro viene spostata di un numero fisso `k` di posti (non necessariamente tre), dove `k` è detto **chiave di cifratura**. 

Scrivere un programma che:
* legga da **standard input** un numero intero `k` (la chiave di cifratura);
* legga da **standard input** un messaggio in chiaro su più righe, terminando la lettura quando, premendo la combinazione di tasti `Ctrl+D`, viene inserito da **standard input** l'indicatore End-Of-File (EOF);
* stampi il messaggio cifrato (ottenuto con chiave di cifratura `k`) corrispondente al messaggio in chiaro letto.

Oltre alla funzione `main()`, il programma deve definire e utilizzare le seguenti funzioni:
* una funzione `LeggiNumero() int` che legge da **standard input** un numero intero e ne restituisce il valore;
* una funzione `LeggiTesto() string` che legge da **standard input** un testo su più righe e terminato dall'indicatore EOF, restituendo un valore `string` in cui è memorizzato il testo letto;
* una funzione `CifraCarattere(c rune, chiave int) rune` che riceve in input un valore `rune` nel parametro `c` ed un valore `int` nel parametro `chiave`, e restituisce un valore `rune` uguale a `c` nel caso in cui `c` non sia una lettera dell'alfabeto inglese, uguale al valore cifrato corrispondente a `c` (ottenuto con chiave di cifratura `chiave`) altrimenti. In particolare, il valore cifrato deve essere minuscolo se `c` è minuscolo e maiuscolo se `c` è maiuscolo;
* una funzione `CifraTesto(s string, chiave int) string` che riceve in input un valore `string` nel parametro `s` ed un valore `int` nel parametro `chiave`, e restituisce un valore `string` ottenuto cifrando ogni carattere di `s` tramite la funzione `CifraCarattere()`.

##### Esempio d'esecuzione:
 
 ```text
$ go run cifrario_cesare.go 
Inserisci un numero: 1
Inserisci un testo su più righe (termina con CTRL D):
Testo di esempio
diviso su righe diverse
Testo cifrato:
Uftup ej ftfnqjp
ejwjtp tv qsjhif ejwfstf

$ go run cifrario_cesare.go
Inserisci un numero: -2
Inserisci un testo su più righe (termina con CTRL D):
AbC

dEf
Testo cifrato:
YzA

bCd
```
# Date (1)

Scrivere un programma che:
* legga da **standard input** un testo su più righe;
* termini la lettura quando viene inserita da standard input una riga vuota (`""`).

Ogni riga di testo è una una stringa senza spazi che codifica una data in uno dei seguenti possibili formati:

1. aaaa/m/g
2. aaaa/m/gg
3. aaaa/mm/g
4. aaaa/mm/gg

Una volta terminata la fase di lettura, il programma deve stampare la codifica nel formato aaaa-mm-gg delle date lette.

Oltre alla funzione `main()` devono essere definite ed utilizzate almeno le seguenti funzioni:
* una funzione `LeggiTesto() []string` che legge da **standard input** un testo su più righe e terminato da una riga vuota (`""`), restituendo un valore `[]string` in cui sono memorizzate le righe del testo letto;
* una funzione `Formatta(s string) string` che riceve in input un valore `string` nel parametro `s` in cui è codificata una data nel formato **aaaa/m/g**, **aaaa/m/gg**, **aaaa/mm/g** o **aaaa/mm/gg**, e restituisce un valore `string` in cui la data in input è codificata nel formato **aaaa-mm-gg**.

##### Esempio d'esecuzione:

```text
$ cat test
2019/04/03
2019/6/4
2019/07/5
2019/4/05
$ go run date.go < test
2019-04-03
2019-06-04
2019-07-05
2019-04-05
```
# Filtra testo

Scrivere un programma che: 
* legga da **standard input** un testo su più righe (alcune delle quali possono essere delle righe vuote (`""`));
* termini la lettura quando, premendo la combinazione di tasti `Ctrl+D`, viene inserito da **standard input** l'indicatore End-Of-File (EOF);
* ristampi le righe del testo letto in cui compaiono cifre.

Oltre alla funzione `main()`, devono essere definite ed utilizzate almeno le seguenti funzioni:
* una funzione `LeggiTesto() []string` che legge da **standard input** un testo su più righe (alcune delle quali possono essere delle righe vuote (`""`)) e terminato dall'indicatore EOF, restituendo un valore `[]string` in cui sono memorizzate le stringhe corrispondenti alle righe del testo letto;
* una funzione `ContieneCifre(s string) bool` che riceve in input un valore `string` nel parametro `s` e restituisce `true` se almeno un carattere in `s` è una cifra, `false` altrimenti;
* una funzione `FiltraTesto(sl []string) []string` che riceve in input un valore `[]string` nel parametro `sl` e restituisce un valore `[]string` in cui sono  memorizzate le stringhe di `sl` in cui compaiono cifre; la funzione deve utilizzare la funzione `ContieneCifre()`.

*Suggerimento:* per sapere se un carattere è una cifra si può far riferimento alla documentazione della funzione `unicode.IsDigit` del package `unicode`.

```text
$ go run filtra_testo.go
riga senza cifre
riga con 1 cifra         
2a riga con 2 cifre
nessuna cifra 
Testo filtrato:
riga con 1 cifra
2a riga con 2 cifre
```
# Date (2)

Scrivere un programma che:
* legga da **standard input** un testo su più righe;
* termini la lettura quando viene inserita da standard input una riga vuota ( "" ).

Ogni riga di testo è una una stringa senza spazi che codifica una data in uno dei seguenti possibili formati:

1. aaaa/m/g
2. aaaa/m/gg
3. aaaa/mm/g
4. aaaa/mm/gg
5. g/m/aaaa
6. gg/m/aaaa
7. g/mm/aaaa
8. gg/mm/aaaa

Una volta terminata la fase di lettura, il programma deve stampare la codifica nel formato aaaa-mm-gg delle date lette. In particolare, le date devono essere stampate in ordine cronologico (dalla più antica alla più recente).

Oltre alla funzione `main()` devono essere definite ed utilizzate almeno le seguenti funzioni:
* una funzione `daInvertire(s string) bool` che riceve in input un valore `string` nel parametro `s` in cui è codificata una data nel formato **aaaa/m/g**, **aaaa/m/gg**, **aaaa/mm/g**, **aaaa/mm/gg**, **g/m/aaaa**, **gg/m/aaaa**, **g/mm/aaaa**, o **gg/mm/aaaa** e restituisce `true` se in `s` è codificata una data nel formato **g/m/aaaa**, **gg/m/aaaa**, **g/mm/aaaa**, o **gg/mm/aaaa**, `false` altrimenti.
* una funzione `Inverti(s string) string` che riceve in input un valore  `string` nel parametro `s` e restituisce un valore `string` in cui il primo carattere è l'ultimo che definisce `s`, il secondo carattere è il penultimo che definisce `s`, ...  e l'ultimo carattere è il primo che definisce `s`.

*Suggerimento:* Una volta codificate nel formato **aaaa-mm-gg**, le date possono essere ordinate cronologicamente utilizzando la funzione `sort.Strings(a []string)` del package `sort`.

##### Esempio d'esecuzione:

```text
$ cat test
2019/04/03
2019/6/4
2019/07/5
2019/4/05
5/5/2019
05/4/2019
7/05/2019
07/09/2019
07/09/2018

$ go run date.go < test      
2018-09-07
2019-04-03
2019-04-05
2019-04-05
2019-05-05
2019-05-07
2019-06-04
2019-07-05
2019-09-07
```
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
