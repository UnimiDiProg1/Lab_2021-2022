# Esercizio filtro

Scrivere un programma che legga da **riga di comando** un numero naturale positivo e, come mostrato nell'**Esempio di esecuzione**, stampi a video ogni cifra del numero inferiore a quella che la segue.

Si assuma un ordinamento delle cifre da sinistra verso destra. Ad esempio, dato il numero 496:
* 4 è la prima cifra del numero; la cifra 4 è seguita dalla cifra 9; 
* 9 è la seconda cifra del numero; la cifra 9 è seguita dalla cifra 6;
* 6 è la terza e ultima cifra del numero; la cifra 6 non è seguita da nessun altra cifra.

Si assuma inoltre che il valore letto da **riga di comando** sia nel formato corretto.

Si noti che l'ultima cifra del numero letto non viene mai stampata.

##### Esempio d'esecuzione:

```text
$ go run esercizio_filtro.go 344
3

$ go run esercizio_filtro.go 34456
345

$ go run esercizio_filtro.go 3

$ go run esercizio_filtro.go 37778
37
```

## Test automatico

L'esercizio filtro è considerato esatto **solo se** rispetta le specifiche date e **solo se** passa il test automatico fornito

##### Inizializzazione del test automatico

Al fine di poter eseguire il test automatico, è prima necessario eseguire **una volta sola** ed **esclusivamente nella cartella relativa a questo esercizio** il comando:
`go mod init filtro` ottenendo il seguente output:

```text
$ go mod init filtro
go: creating new go.mod: module filtro
go: to add module requirements and sums:
        go mod tidy
```

##### Esecuzione del test automatico

Successivamente sarà possibile eseguire il test automatico utilizzando il comando `go test`.
L'esercizio passa il test automatico se il comando `go test` restituisce `PASS`, come nel seguente output:

```text
$ go test
PASS
ok
```

Invece, nel caso in cui l'output dovesse restituire `FAIL`, come nel seguente esempio, significa che almeno un caso tra quelli riportati nell'esempio d'esecuzione non è stato eseguito in modo corretto, ed il filtro è considerato **errato**.
In tal caso `go test` restituirà anche l'output atteso dal programma e l'output effettivo che il programma produce. Questo dato è utile per capire in cosa il risultato del programma differisce dall'output atteso.

```text
$ go test
--- FAIL: TestFiltro (0.00s)
        esercizio_filtro_test.go:40: 
                ...
```


