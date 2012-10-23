package main

import "fmt"

/*

slice の zero 値は nil です。
空の slice は len cap が 0

*/

func main() {
	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}
}