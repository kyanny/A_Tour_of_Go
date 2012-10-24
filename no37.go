package main

import "fmt"

/*

で、関数はクロージャで、という話。
クロージャなので adder が返す関数 pos neg のなかの sum は互いに独立している。

*/

func adder() func(int) int { // adder の型は func(int) 型? どう読むんだろこれ
	sum := 0
	return func(x int) int { // 無名関数か
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
			)
	}
}