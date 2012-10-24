package main

import (
	"fmt"
	"time" // これもここが初。
)

func main() {
	fmt.Println("when's Satuarday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today+0:
		fmt.Println("Today.")
	case today+1:
		fmt.Println("Tomorrow.")
	case today+2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too bar away.")
	}
}