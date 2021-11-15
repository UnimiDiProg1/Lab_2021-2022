# Triangolo

Scrivere un programma che legga da **standard input** un valore intero e sei valori reali:
- il primo valore è il seme (seed) `s` da utilizzare per inizializzare il generatore di numeri casuali;
- il secondo ed il terzo valore sono il coefficiente angolare `m1` e il termine noto `q1` di una retta `r1: y = m1*x + q1` sul piano cartesiano su cui giace la base `AB` di un triangolo `ABC`;
- il quarto ed il quinto valore sono il coefficiente angolare `m2` e il termine noto `q2` di una retta `r2: y = m2*x + q2` sul piano cartesiano su cui giace il lato `BC` di un triangolo `ABC`;
- il sesto ed il settimo valore sono il coefficiente angolare `m3` e il termine noto `q3` di una retta `r3: y = m3*x + q3` sul piano cartesiano su cui giace il lato `AC` di un triangolo `ABC`.

Una volta terminata la fase di lettura, il programma deve generare per 10 volte una coppia di valori reali `px` e `py` che rappresentano rispettivamente l'ascissa e l'ordinata di un punto sul piano cartesiano e, per ciascun punto:
 1. determinare se, a meno di una costante `EPSILON = 1.0e-9`, il punto sta all'interno o all'esterno del triangolo `ABC`;
 2. stampare a video il relativo messaggio (come mostrato nell'**Esempio di esecuzione**).

I valori `px` e `py` devono essere compresi nell'intervallo [0, 10.0).

Si assuma che:
- i valori introdotti dall'utente che definiscono le rette `r1`, `r2`, ed `r3` permettano di caratterizzare un triangolo `ABC` in cui la base `AB` sia parallela all'asse delle ascisse, ed il vertice `C`, opposto alla base, sia al di sopra della retta `r1` su cui giace la base del triangolo:
 ```markdown        
        C     
        /\
       /  \
      /____\
     A      B 
  ```
- i valori introdotti dall'utente che definiscono le rette `r1`, `r2`, ed `r3` permettano di definire un triangolo `ABC` che contenga almeno un punto con coordinare `px` e `py` compresi nell'intervallo [0, 10.0).

*Suggerimento:* Un punto non appartiene al triangolo se sta al di sotto della retta `r1`, al di sopra della retta `r2`, o al di sopra della retta `r3`.

##### Esempio d'esecuzione:

```markdown
$ go run triangolo.go
Inserisci s: 3
Inserisci m1 e q1: 0 2
Inserisci m2 e q2: -1 15
Inserisci m3 e q3: 1 0

Punto (7.199826688373036,6.526308027999122) - Il punto sta all'interno del triangolo.
Punto (9.419605585830052,7.681370946252233) - Il punto sta all'esterno del triangolo.
Punto (8.935331264200833,2.190716471450932) - Il punto sta all'interno del triangolo.
Punto (4.27633484874847,5.076860623844484) - Il punto sta all'esterno del triangolo.
Punto (3.250879908255019,4.684369767626347) - Il punto sta all'esterno del triangolo.
Punto (3.503080765299786,7.956635723501489) - Il punto sta all'esterno del triangolo.
Punto (8.170357818404945,3.493485636935447) - Il punto sta all'interno del triangolo.
Punto (9.603601912420224,4.76799474008711) - Il punto sta all'interno del triangolo.
Punto (4.994476561730858,7.24266412974525) - Il punto sta all'esterno del triangolo.
Punto (4.009944732483285,4.053135323201041) - Il punto sta all'esterno del triangolo.

$ go run triangolo.go
Inserisci s: 5
Inserisci m1 e q1: 0 1
Inserisci m2 e q2: -1 10
Inserisci m3 e q3: 1 0

Punto (8.038448294975284,5.196192472270505) - Il punto sta all'esterno del triangolo.
Punto (9.79385566880146,5.967488889192919) - Il punto sta all'esterno del triangolo.
Punto (4.793723986112438,7.002671557670805) - Il punto sta all'esterno del triangolo.
Punto (4.689419113522166,8.5793442542724) - Il punto sta all'esterno del triangolo.
Punto (2.9902926479032783,9.943523044874215) - Il punto sta all'esterno del triangolo.
Punto (0.01626037515838685,6.51527635860278) - Il punto sta all'esterno del triangolo.
Punto (5.443642101331307,1.303299350890999) - Il punto sta all'interno del triangolo.
Punto (4.965806911581438,2.9708160500789433) - Il punto sta all'interno del triangolo.
Punto (6.277694594891373,4.120684542231575) - Il punto sta all'esterno del triangolo.
Punto (1.2876742934360954,7.403669183503507) - Il punto sta all'esterno del triangolo.
``` 