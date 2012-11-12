package main

import (
	"fmt"
	"math"
)

// 構造体にメソッドを定義できると言ったな、ありゃ嘘だ。
// 実際は同一パッケージ内で宣言した型ならなんでもメソッドを定義できる。
// 他のパッケージ内で宣言された型や、基本型の場合はメソッドを定義できない。

type MyFloat float64

func (f MyFloat) Abs() float64 { // f はただの float な変数。関数ではないよ。
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2) // math.Sqrt2 は定数。
	fmt.Println(f.Abs())
}