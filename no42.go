package main

import (
	"fmt"
	"time"
)

/*

すごい長い if を書くかわりにこういうのも試していては。

*/

func main() {
	t := time.Now()
	switch { // switch {} ってやると loop {} みたいにできる。条件は省略されると true になる。
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}