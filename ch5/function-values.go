package main

import "fmt"

func square(n int) int     { return n * n }
func negative(n int) int   { return -n }
func product(m, n int) int { return m * n }

func example1() {
	f := square
	value := f(12)

	fmt.Println(value)
}

// This example doesn't work. Even though functions are first-class citizens in go, they're not comparable so can't be used in maps etc
func example2() {
	operators := map[string]func{
		"Square": square
	}
}

func main() {
	example1()
	example2()
}
