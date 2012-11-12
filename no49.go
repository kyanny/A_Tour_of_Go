package main

import (
	"fmt"
	"math"
)

// Go にクラスはないけど構造体にメソッドをつけられる

type Vertex struct {
	X, Y float64
}


// メソッドは func とメソッド名の間に引数リストを書く
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Println(v.Abs())
}