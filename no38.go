package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow { // 配列を range でぶんまわすと for はイテレータになって何番目かも受け取れる
		fmt.Printf("2**%d = %d", i, v)
	}
}