package main

import "fmt"
import "math/cmplx"

// complex number ... 複素数
// cube root ... 立方根、三乗根 (平方根の親戚)
//
// http://go-tour-jp.appspot.com/#49 をみて意味がわかった。
// ニュートン法を素直に実装して (平方根?のときと同様素朴に10回ループ)
// それを cmplx.Pow と見比べろと。
func Cbrt(x complex128) complex128 {
	z := complex128(1)
	for i := 0; i < 10; i++ {
		z = z - (z*z*z - x) / (3 * z*z)
		fmt.Println(z, i)
	}
	return z
}

func main() {
	fmt.Println(Cbrt(2))
	fmt.Println(cmplx.Pow(2, 1.0/3.0))
}