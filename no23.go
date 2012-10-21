package main

import "fmt"

type Vertex struct { // でた struct! C でポインタの次によくわからないシロモノだった...
	X int // 構造体の中のフィールド宣言でも型は後ろに置く
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2}) // Vertext({}) じゃあないのか
//	fmt.Println(Vertex({1, 2})) // やっぱりダメみたいね
}