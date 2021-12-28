# Esercizio 3

Un'utenza di telefonia mobile è identificata dal numero del telefono mobile e da un codice che identifica la scheda SIM che gli corrisponde.

Ne consegue che, in un dato istante temporale, ad un numero di telefono mobile possa corrispondere una ed una sola scheda SIM.

Comunque, nel corso del tempo, ad un numero di telefono mobile possono corrispondere diversi codici relativi a schede SIM. 

L'utenza di telefonia mobile attiva è quella caratterizzata dal codice più recente relativo ad una scheda SIM.

## Parte 1

Scrivere un programma che:
* legga da **standard input** una sequenza di righe di testo;
* termina la lettura quando, premendo la combinazione di tasti `Ctrl+D`, viene inserito da **standard input** l'indicatore End-Of-File (EOF).

Ogni riga di testo è nel formato
```text
NUMERO_TELEFONO;CODICE_SIM
```
dove `NUMERO_TELEFONO` è il numero telefonico (valore composto da sole cifre) di una utenza, mentre `CODICE_SIM` è il codice che identifica la scheda SIM associata (valore composto da caratteri alfanumerici).

Ciascuna riga di testo descrive quindi un'utenza di telefonia mobile.
Le utenze di telefonia mobile sono specificate in ordine cronologico:
 la prima riga descrive l'utenza di telefonia mobile meno recente, l'ultima riga descrive l'utenza di telefonia mobile più recente.

Definire la struttura `Utenza` per memorizzare il numero telefonico e il codice della SIM associata.

Implementare le funzioni:
* `LeggiUtenze() (utenze []Utenza)` che:
    1. legge da **standard input** una sequenza di righe di testo, terminando la lettura quando viene letto l'indicatore End-Of-File (EOF);
    2. restituisce un valore `[]Utenza` nella variabile `utenze` in cui è memorizzata la sequenza di istanze del tipo `Utenza` inizializzate con i valori letti da **standard input**.
    
## Parte 2

Il registro telefonico di un operatore è la lista di tutte le utenze telefoniche passate (non attive) e presenti (attive).
Il registro contiene, per ogni numero di telefono, l'elenco di utenze associate in ordine cronologico.
Definire il tipo di dato `RegistroTelefonico` (alias di `map[string][]string`) che associa ad ogni numero di telefono tutte le SIM in ordine cronologico.

Implementare le funzioni:
* `InizializzaRegistro() (registro RegistroTelefonico)` che inizializza un'istanza di tipo `RegistroTelefonico` (un registro telefonico vuoto) e la restituisce all'interno della variabile `registro`;
* `AggiungiUtenza(registro RegistroTelefonico, utenza Utenza) (registroAggiornato RegistroTelefonico)` che riceve in input un'istanza di tipo `RegistroTelefonico` nella variabile `registro` e un'istanza di tipo `Utenza` nella variabile `utenza`. Se il numero dell'utenza non è presente nel registro, viene aggiunta una nuova voce. Altrimenti, viene solamente aggiunta la nuova sim al numero di telefono.  La funzione restituisce il registro telefonico aggiornato nella variabile `registroAggiornato`. 
* `Identifica(registro RegistroTelefonico, telefono string) (codiceSIM string)` che riceve in input un'istanza di tipo `RegistroTelefonico` nella variabile `registro` e un valore di tipo `string` nella variabile `telefono` (un numero di telefono mobile).
La funzione restituisce il codice della SIM (presente in `registro`) corrispondente a `telefono` nella variabile `codiceSIM`.
Se in `registro` sono presenti più codici SIM corrispondenti a `telefono`, viene restituito il codice più recente, ovvero quello che nella sequenza di input è letto per ultimo.
Se in `registro` non è presente alcun codice SIM corrispondente a `telefono`, viene restituito il valore `""`.

Scivere un programma che:
* Richiami la funzione `LeggiUtenze` per leggere da standard input una sequenza di utenze
* Usi le funzioni `InizializzaRegistro` e `AggiungiUtenza` per inizializzare il registro ed inserirvi ogni utenza letta da standard input
* Stampi per ogni numero di telefono nel registro che inizia con 338, la SIM più recente restituita dalla funzione `Identifica`.

##### Esempio d'esecuzione:

Nota bene: siccome l'output del programma è lungo, vengono riportate solo le prime tre righe e le ultime tre righe separate da `[...]` per brevità. 

```text
$ go run esercizio_3.go < registro1.txt 
Il numero 338245567 è associato alla sim 464c1f62ff7f8dafff0a20e6bb2fa20e
Il numero 338345567 è associato alla sim 5fa42ef883cf1f749b55c213dda314ab
Il numero 338345467 è associato alla sim 809eb5802ccf66708614f75cd8800bc7
[...]
Il numero 338145167 è associato alla sim edaf427ea16089fc1892eccfd39b94b6
Il numero 338245067 è associato alla sim 23e8e8d3cff934d8302db861b3d4a14e
Il numero 338245467 è associato alla sim 5e1d885e4e462d556f807bbd29f3e62a

$ go run esercizio_3.go < registro2.txt 
Il numero 338445367 è associato alla sim e7e818dd24e61103189fc39a73e98c4d
Il numero 338445267 è associato alla sim 2babaeb631887cd90b82e56e1fcdba7b
Il numero 338345467 è associato alla sim 7514c20e15d760dc5f69e9aefe0784d8
[...]
Il numero 338345267 è associato alla sim f2de3e746944aa8fd1798e4c97baf751
Il numero 338145167 è associato alla sim 202620525af16637f8551c5d1407bb9d
Il numero 338545567 è associato alla sim 9556aa951e6a1e5315fa0d65addfa5cb
```