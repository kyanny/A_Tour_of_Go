package main

import (
	"math/cmplx" // complex??
	"fmt"
)

var ( // var も ()　で複数宣言可能なのか
ToBe bool = false
MaxInt uint64 = 1<<64 - 1 // << ビットシフト? こういう例 C っぽい
z complex128 = cmplx.Sqrt(-1+12i) // complex て複素数を含む型?
) // var () はインデントがうまくいかない

func main() {
	const f = "%T(%v)\n" // %T で型名を表示するのか
	fmt.Printf(f, ToBe, ToBe) // そして書式付きプリントなので Println じゃなく Printf
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)
}
