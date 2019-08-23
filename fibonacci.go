// fibonacci - prints the first N numbers of the sequence
package main

import "fmt"

// fibonacci returns a function which returns the next number (int)
// in the Fibonacci sequence.
//
// It uses closure to track the next fibonacci number in the sequence.
func fibonacci() func() int {
	var fib = map[int]int{0: 0, 1: 1, 2: 1}
	var index int

	return func() int {
		_, ok := fib[index]
		if !ok {
			fib[index] = fib[index-1] + fib[index-2]
		}
		//v = fib[index]
		index++
		return fib[index-1]
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
