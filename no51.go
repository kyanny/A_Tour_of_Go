package main

import (
	"fmt"
	"math"
)

/*

メソッドをポインタ型に定義するのと型そのものに定義するのと二種類ある。

ポインタ型を使うメリットは二つ。

 1. 値そのものはコピーしないのでパフォーマンスが良い
 2. 呼び出し元の値をいりじたい（影響を与えたい）ときは値をコピーしてもダメなのでポインタ型が必要

Scale2 Abs2 はそれぞれ Vertex 型そのもののmethodなので値を受け取って値を返している。
なので受け取った構造体をいじっても呼び出し元の構造体には影響を与えない。

*/

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vertex) Scale2(f float64) {
	v.X = v.X * f
  v.Y = v.Y * f
}

func (v Vertex) Abs2() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	v.Scale(5)
	fmt.Println(v, v.Abs())

	v2 := &Vertex{3, 4}
	v2.Scale2(5)
	fmt.Println(v2, v2.Abs2())
}