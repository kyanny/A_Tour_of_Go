package main

import "fmt"

type Vertex struct { // vertex は頂点とかいう単語
	X int
	Y int
}

func main() {
	p := Vertex{1, 2}
	q := &p // でた!! ポインタ!! Go はポインタがあるけどポインタ演算 arithmetic は無いらしい。
	q.X = 1e9 // 構造体にはポインタ (参照?) を通じてもアクセス可能らしい。これ何のためにあるのか今はまだよくわからないな。
	fmt.Println(p)
}