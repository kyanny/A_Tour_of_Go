package main

import "fmt"

type Vertex struct {
	X, Y int
}

func main() {
	v := new(Vertex)
	fmt.Println(v)
	v.X, v.Y = 11, 9
	fmt.Println(v)
}

/*

 new(T) という関数は構造体の各フィールドの値を 0 で初期化した状態で、その初期化済み構造体へのポインタを返す。??

 var t *T = new(T)
 t := new(T) // t has type *T

 ...ということらしいけど...型難しい...

*/