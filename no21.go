package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10), // あーそうか、こっちは else で else では return してないので、 >= の printf が印字されたあと、次の行で lim も印字されるのか。
		pow(3, 3, 20),
		 )

	fmt.Println("----")
	fmt.Println(pow(3, 2, 10)) // 9
	fmt.Println("----")
	fmt.Println(pow(3, 3, 20)) // 27 >= 20\n20
}