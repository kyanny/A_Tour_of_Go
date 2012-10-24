package main

import "fmt"

func main() {
	pow := make([]int, 10) // 復習: これは slice を作ってて要素上限が 10 個だっけ?
  for i := range pow { // _, value でうけなければインデックスだけ得られる
		pow[i] = 1<<uint(i)
	}
	for _, value := range pow { // _ で受けることで値を捨てる (_ が予約語的なものなのか、単に不要な変数を受ける仮の名前というイディオムなのか?)
		fmt.Printf("%d\n", value)
	}
}