package main

import (
	"fmt"
	"math"
)

func main() {
	hypot := func(x, y float64) float64 { // 関数も変数で、という話。 JavaScript を知ってればわかる。
		return math.Sqrt(x*x + y*y) // return 必須なのだろうか、めんどい
	}

	fmt.Println(hypot(3, 4))
}