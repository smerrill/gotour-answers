package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	x := make([]int, 1000)
	x[0] = 1
	x[1] = 1
	index := 0

	return func() int {
		if index > 1 {
			x[index] = x[index-1] + x[index-2]
		}
		output := x[index]
		index++
		return output
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
