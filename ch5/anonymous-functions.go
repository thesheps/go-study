package main

import "fmt"

func squares() func() int {
	var x int

	return func() int {
		x++
		return x * x
	}
}

func main() {
	val := squares()

	fmt.Println(val())
	fmt.Println(val())
	fmt.Println(val())
}
