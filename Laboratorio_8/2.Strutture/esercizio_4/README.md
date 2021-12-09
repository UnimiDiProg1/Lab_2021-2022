# Rubrica (1)

Si consideri una rubrica in cui:
* Ogni contatto deve avere un cognome, un nome, un indirizzo ed un numero di telefono.
* Ogni indirizzo deve contenere le seguenti informazioni: via, numero civico, CAP e città.
* Non possono esistere due contatti con lo stesso cognome e lo stesso nome.

## Strutture dati da definire

Definire i seguenti tipi di dati:

* `Indirizzo`: una struttura che memorizzi un indirizzo nei seguenti campi: `via`, `cap` e `città` di tipo `string` e `numeroCivico` di tipo `uint`;

* `Contatto`: una struttura che memorizzi i dati di un contatto nei seguenti campi: `cognome`, `nome` e `telefono` di tipo `string` e `indirizzo` di tipo `Indirizzo`;

* `Rubrica`: una slice dove ogni elemento è di tipo `Contatto`. Ogni elemento della slice rappresenta un contatto inserito nella rubrica. Una slice vuota rappresenta una rubrica vuota. 

## Funzioni da implementare

Implementare le funzioni:

* `NuovaRubrica() Rubrica` che restituisce un valore `Rubrica` rappresentante una rubrica senza alcun contatto inserito;

* `InserisciContatto(r Rubrica, cognome, nome string, via string, numero uint, cap, città string, telefono string) Rubrica` che data la rubrica `r` restituisce una nuova rubrica in cui è inserito il nuovo contatto creato con i dati passati come argomento. Se la rubrica `r` contiene già un contatto con identici `nome` e `cognome` non avviene nessuna modifica e la rubrica restituita è `r` stessa;
  
* `EliminaContatto(r Rubrica, cognome, nome string) Rubrica` che
  restituisce una rubrica in cui è eliminato il contatto avente `nome` e `cognome` uguali a quelli passati in input. Se tale contatto non esiste, la rubrica restituita è `r` stessa;
  
* `StampaContatto(c Contatto)` che stampa a video i dettagli del contatto `c` nel seguente formato: `nome cognome: via numeroCivico, città, cap - Tel. telefono\n`;

* `StampaRubrica(r Rubrica)` che stampa a video tutti i contatti presenti nella rubrica utilizzando per ogni contatto la funzione `StampaContatto()`. La stampa deve rispettare il seguente formato:
  ```text
  Rubrica:\n
  [1] - nome cognome: via numeroCivico, città, cap - Tel. telefono\n
  [2] - nome cognome: via numeroCivico, città, cap - Tel. telefono\n
  [3] - nome cognome: via numeroCivico, città, cap - Tel. telefono\n
  ```
  
* `AggiornaContatto(rubrica Rubrica, cognome, nome string, via string, numero uint, cap, città string, telefono string) Rubrica` che aggiorna i dettagli del contatto identificato da `nome` e `cognome` e restituisce la rubrica con il contatto aggiornato. Se il contatto non esiste, viene restituita la rubrica `r` stessa.

##### Esempio d'esecuzione:

Il formato dell'output dopo alcune operazioni di inserimento, cancellazione e modifica della rubrica dovrebbe risultare simile al seguente:
```text
$ go run rubrica.go
Rubrica:
[1] - Mario Rossi: Via Festa del Perdono 11, 20122, Milano - Tel. 02503111
[2] - Anna Rossi: Via Festa del Perdono 11, 20122, Milano - Tel. 02503111
[3] - Carlo Rossi: Via Festa del Perdono 11, 20122, Milano - Tel. 02503111
Rubrica:
[1] - Anna Rossi: Via Festa del Perdono 11, 20122, Milano - Tel. 02503111
[2] - Carlo Rossi: Via Festa del Perdono 11, 20122, Milano - Tel. 02503111
Rubrica:
[1] - Anna Rossi: Via S. Sofia 25, Milano, 20122 - Tel. 
[2] - Carlo Rossi: Via Festa del Perdono 11, 20122, Milano - Tel. 02503111
```