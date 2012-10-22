package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex // map がありますよ、という話。

func main() {
	m = make(map[string]Vertex) // map は new じゃなくて make で初期化する。 nil map is empty and ... よくわからず。必ず値を入れて初期化する必要があるってこと?
	m["Bell Labs"] = Vertex{
		40.68433, 74.39967,
	}
	fmt.Println(m["Bell Labs"])
}