package main

import (
	"fmt"
)

func main() {
	count := 0
	for i := 1000; i <= 9999; i++ {
		for j := i; j <= 9999; j++ { // Not count twice
			if isEvenNumber(i * j) {
				count++
			}
		}
	}

	fmt.Printf("%v", count)
}

func isEvenNumber(num int) bool {
	str := fmt.Sprintf("%v", num)
	if str[0] == str[len(str)-1] {
		return true
	}
	return false
}
