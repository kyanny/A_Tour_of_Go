package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{ // map てのは連想配列のことか。リテラル表記もできる。 struct と似てるけどキーが必須。
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408, // (2) このカンマも気持ち悪いけどね... こちらも必須ぽい。
	}, // (1) このカンマがとても気持ち悪い... 必須みたい。ないとエラー。
}

func main() {
	fmt.Println(m)
}