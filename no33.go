package main

import "fmt"

/*

引き続き slice

 s[lo:hi] は lo 番目から hi-1 番目までの項目を含む新しい slice を返す。
 s[lo:lo] は 0 個、 s[lo:lo+1] は 1 個の要素。

*/

func main() {
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p == ", p)
	fmt.Println("p[1:4] == ", p[1:4]) // slice から添字で slice を取り出すには [n:m] を使う。 python ぽい。
	
	// missing low index implies 0
	fmt.Println("p[:3] ==", p[:3]) // [:m] は [0:m] と同じ

	// missing high index implies len(s)
	fmt.Println("p[4:] ==", p[4:]) // [n:] は [n:len()] と同じ
}