# Valori rappresentabili

Il seguente programma stampa a video i limiti dei valori rappresentabili con i diversi tipi di dato numerici.
Per ogni tipo di dato, i corrispondenti limiti sono definiti come costanti nel package `math`.
Si stampi a video il valore di tali limiti eseguendo il programma.

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("uint8:", 0, math.MaxUint8)
	fmt.Println("uint16:", 0, math.MaxUint16)
	fmt.Println("uint32:", 0, math.MaxUint32)
	fmt.Println("uint64:", 0, uint64(math.MaxUint64))

	fmt.Println("int8:", math.MinInt8, math.MaxInt8)
	fmt.Println("int16:", math.MinInt16, math.MaxInt16)
	fmt.Println("int32:", math.MinInt32, math.MaxInt32)
	fmt.Println("int64:", math.MinInt64, math.MaxInt64)

    /*  Floating-point limit values: 
        - Max is the largest finite value representable by the type.
        - SmallestNonzero is the smallest positive, non-zero value
        representable by the type. */

	fmt.Println("float32 - SmallestNonzero:", math.SmallestNonzeroFloat32)
	fmt.Println("float32:", -math.MaxFloat32, math.MaxFloat32)
	fmt.Println("float64 - SmallestNonzero:", math.SmallestNonzeroFloat64)
	fmt.Println("float64:", -math.MaxFloat64, math.MaxFloat64)

	fmt.Println("complex64:", complex(-math.MaxFloat32, -math.MaxFloat32), 
        complex(math.MaxFloat32, math.MaxFloat32))
	fmt.Println("complex128:", complex(-math.MaxFloat64, -math.MaxFloat64), 
        complex(math.MaxFloat64, math.MaxFloat64))

}
```