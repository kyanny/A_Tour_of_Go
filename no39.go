package main

import "fmt"

func main() {
	pow := make([]int, 10) // 復習: これは slice を作ってて要素上限が 10 個だっけ?
	for i := range pow {
		pow[i] = 1<<uint(i)
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}