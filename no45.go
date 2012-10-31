package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
	outer_slice := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		inner_slice := make([]uint8, dx)
		for x := 0; x < dx; x++ {
			inner_slice[x] = uint8(x)
		}
		outer_slice[y] = inner_slice
	}
	return outer_slice
}

func main() {
	pic.Show(Pic)
}