package main

import "fmt"

/*

slice は make で作ることができる。 zeroed array とか書いてあるのでやっぱり slice と別の array もあるってことなんだろうか。よくわからない。

 a := make([]int, 5) と書く。

閉じカッコの位置が紛らわしい。 int の後ろにはない。これは要素数 5 の slice を作る。 len(a) = 5

slice には capacity という制限値を与えることができる。

 b := make([]int, 0, 5)

これで要素数 0 個で最大要素数が 5 個の slice を作る。 cap(b) で capacity が得られる。

 b = b[:cap(b)] とか re-slice すると cap までの範囲で slice の要素数を大きくできる。
 b = b[1:] で要素一個減る、これは cap もなぜか一個減る。

*/

func main() {
	a := make([]int, 5)
	printSlice("a", a)
	b := make([]int, 0, 5)
	printSlice("b", b)
	c := b[:2]
	printSlice("c", c)
	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) { // 仮引数が slice の場合、型は []int と書く。
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}