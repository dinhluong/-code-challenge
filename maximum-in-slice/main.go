package main

import (
	"fmt"
)

func main() {
	slices := []int{10, 20, 39, 30, 28, 38, 72, 18}
	// Find max here
	max := slices[0]
	for _, sl := range slices[1:] {
		if sl > max {
			max = sl
		}
	}
	fmt.Println(max)
}
