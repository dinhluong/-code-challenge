package main

import (
	"fmt"
)

func main() {

	for i := 0; i <= 20; i++ {
		if i%3 == 0 {
			if i%5 == 0 {
				fmt.Printf("fizzbuzz\n")
				continue
			} else {
				fmt.Printf("fizz\n")
				continue
			}
		}
		if i%5 == 0 {
			fmt.Printf("buzz\n")
			continue
		}
		fmt.Printf("%v\n", i)
		// fmt.Printf("%v", i)

	}
	// time.Sleep(10 * time.Second)
}
