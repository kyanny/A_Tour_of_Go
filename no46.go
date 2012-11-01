package main

import "fmt"

func fibonacci() func() int {
	n := 0
	m := 1
	return func() int {
		ret := n
		n = m
		m = ret + m
		return ret
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}