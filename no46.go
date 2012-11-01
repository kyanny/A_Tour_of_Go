package main

import "fmt"

func fibonacci() func() int {
	n := 0
	m := 1
	return func() int {
		retval := n
		n, m = m, n + m
		return retval
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}