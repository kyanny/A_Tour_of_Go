package main

import "fmt"

func fibonacci() func(int) int {
	var f func(n int) int
	f = func(n int) int {
		x := 0
		if n == 0 {
			x = 0
		} else if n == 1 {
			x = 1
		} else {
			x = f(n-1) + f(n-2)
		}
		return x
	}
	return f
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}