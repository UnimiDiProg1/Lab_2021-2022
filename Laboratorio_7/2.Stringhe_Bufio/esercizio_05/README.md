# Statistiche testo (2)

Estendere il programma precedente in modo tale che il conteggio della lunghezza media delle parole non prenda in considerazione numeri e caratteri di separazione.
Si implementi una funzione `contaLettere(string) int` che, data in input una parola derivata dal testo, restituisca il numero effettivo di lettere.


*Suggerimento:* per sapere se un carattere rappresenta una lettera utilizza la funzione `unicode.IsLetter()` del package `unicode`. Utilizzare `go doc unicode.IsLetter` per scoprire il suo funzionamento. 


##### Esempio d'esecuzione:

```text
$ go run statistiche.go    
Inserisci un testo su più righe (termina con Ctrl+D):
Testo di prova...
su cui provare le statistiche!
Statistiche:
Numero linee: 2
Numero parole: 8
Lunghezza media: 4.625
```
