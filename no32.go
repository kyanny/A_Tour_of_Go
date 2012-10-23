package main

import "fmt"

/*

Slice ってのがあって、えーと Array はでてきてないんだっけ？
要するに Array のこと。 []T で T 型の値の配列、という型。

*/

func main() {
	p := []int{2, 3, 5, 7, 11, 13} // slice リテラル書くときも {} ブレースを使うのが go 流っぽい。
	fmt.Println("p ==", p)

	for i := 0; i < len(p); i++ { // len() で slice の長さを得る
		fmt.Printf("p[%d] == %d\n", i, p[i])
	}
}