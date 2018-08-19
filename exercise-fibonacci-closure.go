package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int
func fibonacci() func() int {
	first := 0
	second := 1

	return func() int {
		current := first
		first = second
		second = current + first

		return current
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
