# Rubrica (1)

Si consideri la rubrica dell'esercizio **3 Rubrica (1)**.

Si modifichi il programma in modo tale che le operazioni sulla rubrica vengano gestite tramite dei comandi letti da **standard input**.

Il programma deve:
* creare una rubrica vuota;
* leggere da **standard input** un testo su più righe (ogni riga è un comando);
* terminare la lettura quando viene inserito da standard input l'indicatore EOF (CTRL+D);
* eseguire in sequenza i comando letti eseguendo le funzioni corrispondenti. 

Ogni comando è una riga di testo in cui il primo carattere determina il tipo dell'operazione:
* `I`: inserimento di un contatto
* `E`: cancellazione di un contatto
* `S`: stampa della rubrica
* `A`: aggiornamento di un contatto

I valori successivi rappresentano i valori da assegnare agli argomenti delle funzioni che effettuano l'operazione richiesta (le funzioni implementate nell'esercizio **3 Rubrica (1)**. I valori sono sempre separati tra di loro dal carattere `;`.

### Inserimento

Il comando per l'inserimento di un nuovo contatto ha il seguente formato:
```text
I;cognome;nome;via;numeroCivico;cap;città;telefono
```

### Cancellazione

Il comando per la cancellazione di un nuovo contatto ha il seguente formato:
```text
E;cognome;nome
```

### Stampa

Il comando per la stampa della rubrica è il seguente:
```text
S
```

### Aggiornamento di un contatto

```text
A;cognome;nome;via;numeroCivico;cap;città;telefono
```

*Suggerimento:* per separare i valori presenti all'interno di una riga di testo utilizzate la funzione `strings.Split()` (usate `go doc strings.Split` per leggerne la documentazione).

##### Esempio d'esecuzione:
```text
$ cat comandi.txt 
I;Mario;Rossi;Via Celoria;18;20122;Milano;02503111
I;Mario;Rossi;Via Celoria;18;20122;Milano;
S
I;Elena;Bianchi;Via Celoria;18;20122;Milano;02503111
S
E;Mario;Rossi
S
A;Elena;Bianchi;Via Festa del perdono;7;20122;Milano;02503111
S

$ go run rubrica.go < comandi.txt 
Rubrica:
[1] - Mario Rossi: Via Celoria 18, 20122, Milano - Tel. 02503111
Rubrica:
[1] - Mario Rossi: Via Celoria 18, 20122, Milano - Tel. 02503111
[2] - Elena Bianchi: Via Celoria 18, 20122, Milano - Tel. 02503111
Rubrica:
[1] - Elena Bianchi: Via Celoria 18, 20122, Milano - Tel. 02503111
Rubrica:
[1] - Elena Bianchi: Via Festa del perdono 7, Milano, 20122 - Tel. 02503111
```