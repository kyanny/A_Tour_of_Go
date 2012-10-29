package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (float64, int) {
	z := float64(0)
	for i := 0; i < 10; i++ {
		z = z - ( (math.Pow(z, 2) - x) / math.Pow(2, z))
		if math.Abs(math.Sqrt(x) - z) < 0.01 {
			return z, i
		}
	}
	return z, 10
}

func main() {
	z1, i1 := Sqrt(1)
	fmt.Printf("z = %f (loop#%d)\n", z1, i1)
	fmt.Printf("math.Sqrt(1) = %f\n", math.Sqrt(1))
	fmt.Printf("delta = %f\n\n", math.Abs(z1 - math.Sqrt(1)))

	z2, i2 := Sqrt(2)
	fmt.Printf("z = %f (loop#%d)\n", z2, i2)
	fmt.Printf("math.Sqrt(2) = %f\n", math.Sqrt(2))
	fmt.Printf("delta = %f\n\n", math.Abs(z2 - math.Sqrt(2)))

	z3, i3 := Sqrt(3)
	fmt.Printf("z = %f (loop#%d)\n", z3, i3)
	fmt.Printf("math.Sqrt(3) = %f\n", math.Sqrt(3))
	fmt.Printf("delta = %f\n\n", math.Abs(z3 - math.Sqrt(3)))

	z5, i5 := Sqrt(5)
	fmt.Printf("z = %f (loop#%d)\n", z5, i5)
	fmt.Printf("math.Sqrt(5) = %f\n", math.Sqrt(5))
	fmt.Printf("delta = %f\n\n", math.Abs(z5 - math.Sqrt(5)))

	z7, i7 := Sqrt(7)
	fmt.Printf("z = %f (loop#%d)\n", z7, i7)
	fmt.Printf("math.Sqrt(7) = %f\n", math.Sqrt(7))
	fmt.Printf("delta = %f\n\n", math.Abs(z7 - math.Sqrt(7)))
}