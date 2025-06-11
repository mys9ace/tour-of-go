package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool                              // Default value is false
	MaxInt uint64     = 1<<64 - 1            // Maximum value for uint64
	z      complex128 = cmplx.Sqrt(-5 + 12i) // Complex number with real and imaginary parts
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)     // Prints: Type: bool Value: false
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt) // Prints: Type: uint64 Value: 18446744073709551615
	fmt.Printf("Type: %T Value: %v\n", z, z)           // Prints: Type: complex128 Value: (3+2i)
}
