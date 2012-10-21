package main

import "fmt"

type Vertex struct { // そもそも type 構造体名 struct { } という書式をちゃんと覚えないと...
	X int
 	Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 4 // 構造体のフィールドには . (dot) でアクセスできる。参照代入ともに可能。
	fmt.Println(v.X)
}