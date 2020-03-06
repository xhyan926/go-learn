package main

import (
	"fmt"
)

func Sqrt(x float64) {

	z := 2.0
	for i := 0; i < 100; i++ {
		z -= (z*z - x) / (2 * x)
		fmt.Println(z)
	}
}

func main() {
	Sqrt(2)
}
