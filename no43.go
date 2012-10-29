package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(0)
	for i := 0; i < 10; i++ {
		z = z - ( (math.Pow(z, 2) - x) / math.Pow(2, z))
	}
	return z
}

func main() {
	fmt.Println(Sqrt(1))
	fmt.Println(math.Sqrt(1))
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
	fmt.Println(Sqrt(3))
	fmt.Println(math.Sqrt(3))
}