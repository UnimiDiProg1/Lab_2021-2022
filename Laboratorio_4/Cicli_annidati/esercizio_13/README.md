# Triangoli

Scrivere un programma che legga da **standard input** un intero `n > 1` e stampi, utilizzando il carattere `*`, il perimetro di due triangoli rettangoli con base e altezza di lunghezza `n`, posizionati come mostrato nell'**Esempio d'esecuzione**.
 
##### Esempio d'esecuzione:

```text
$ go run triangoli.go
2
**
 *
 * 
 **

$ go run triangoli.go
3
***
 **
  *
  *  
  ** 
  ***

$ go run triangoli.go
4
****
 * *
  **
   *
   *   
   **  
   * * 
   ****

$ go run triangoli.go 6
******
 *   *
  *  *
   * *
    **
     *
     *     
     **    
     * *   
     *  *  
     *   * 
     ******
```