# Cerchio

Scrivere un programma che legga da **standard input** un valore intero e tre valori reali:
- il primo valore è il seme (seed) `s` da utilizzare per inizializzare il generatore di numeri casuali;
- il secondo ed il terzo valore, `xC` e `yC`, rappresentano rispettivamente l'ascissa e l'ordinata di un punto `C` sul piano cartesiano: il centro di un cerchio;
- il quarto valore, `raggio`, è il raggio del cerchio con centro `C`.

Una volta terminata la fase di lettura, il programma deve generare per `1.000.000` di volte una coppia di valori reali `px` e `py` che rappresentano rispettivamente l'ascissa e l'ordinata di un punto sul piano cartesiano e, come mostrato nell'**Esempio di esecuzione**: 
1. per ogni punto che, a meno di una costante `EPSILON = 1.0e-6`, giace sulla circoferenza del cerchio con centro `C`, deve stampare a video il relativo messaggio;
2. deve stampare l'ascissa e l'ordinata del punto che, di almeno una costante `EPSILON = 1.0e-6`:

    a) è all'esterno del cerchio;

    b) ha distanza massima dal centro `C`;

3. deve stampare l'ascissa e l'ordinata del punto che, di almeno una costante `EPSILON = 1.0e-6`:

    a) è all'interno al cerchio;

    b) ha distanza minima dal centro `C`;

I valori `px` e `py` devono essere compresi rispettivamente negli intervalli `[xC-raggio, xC+raggio)` e `[yC-raggio, yC+raggio)`.

##### Esempio d'esecuzione:

```markdown
$ go run cerchio.go
Inserisci s: 1
Inserisci i valori dell'ascissa e dell'ordinata del punto C (il centro del cerchio): 5 5
Inserisci il valore del raggio: 1

Il punto (4.15954073279876,5.541874673524392) giace sulla circonferenza del cerchio.

Il punto (4.0003279909645535,4.001517285035352) è quello all'esterno del cerchio che ha distanza massima dal centro C.
Distanza: 1.4129090054678466

Il punto (4.999592781833336,5.0000449681220855) è quello all'interno del cerchio che ha distanza minima dal centro C.
Distanza: 0.00040969350405560887

$ go run cerchio.go
Inserisci s: 3
Inserisci i valori dell'ascissa e dell'ordinata del punto C (il centro del cerchio): 4 4
Inserisci il valore del raggio: 0.5

Il punto (4.497366185594927,3.948755997236377) giace sulla circonferenza del cerchio.
Il punto (4.253280394019776,3.568897687168935) giace sulla circonferenza del cerchio.
Il punto (3.9590078930698014,3.501683012145147) giace sulla circonferenza del cerchio.
Il punto (3.8031515161989375,4.4596205489968055) giace sulla circonferenza del cerchio.

Il punto (4.499978695636602,4.4994651760442474) è quello all'esterno del cerchio che ha distanza massima dal centro C.
Distanza: 0.7067136323656061

Il punto (4.0002120210822625,3.9996262973943617) è quello all'interno del cerchio che ha distanza minima dal centro C.
Distanza: 0.0004296586747461331
```