package main

import "fmt"

type Vertex struct { // でた struct! C でポインタの次によくわからないシロモノだった...
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2}) // Vertext({}) じゃあないのか
//	fmt.Println(Vertex({1, 2})) // やっぱりダメみたいね
}