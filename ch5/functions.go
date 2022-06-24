package main

import "fmt"

func fib(term int, sequence []int) int {
	if term < 2 {
		return term
	}

	return fib(term-1) + fib(term-2)
}

func main() {
	value := fib(3, []int{})

	fmt.Println(value)
}
