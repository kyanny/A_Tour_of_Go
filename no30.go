package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{ // map の定義は map[キーの型名]値の型名 {} という感じなのかな。
	"Bell Labs": {40.68433, -74.39967}, // そしてトップレベルの値の型が定義(宣言)の型と同じ場合は型名を省略できる?? 意味取り違えてそう。いやあってるのか。リテラル表記の値のほうの型名を省略できる。
	"Google": {37.42202, -122.08408}, 
}

func main() {
	fmt.Println(m)
}