package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	p = Vertex{1, 2} // has type Vertex
	q = &Vertex{1, 2} // has type *Vertex
	r = Vertex{X: 1} // Y:0 is implicit
	s = Vertex{} // X:0 and Y:0
//	o = Vertex{Y:1, X:3}
)

func main() {
	fmt.Println(p, q, r, s)
//	fmt.Println(o)
}

// q は &{1, 2} という値を持つのか... &{} て意味がわからん。
// Vertex{Y:1, X:2} みたいな書き方もいいみたい。 irrelevant 順番は重要ではない。