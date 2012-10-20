package main

import "fmt"

var x, y, z int = 1, 2, 3 // 型宣言つき
var c, python, java = true, false, "no!" // 型宣言省略できる
// var foo, bar int, baz string = 1, 2, "3" // こういうふうには書けない

func main() {
	fmt.Println(x, y, z, c, python, java)
}