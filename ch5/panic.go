package main

import "fmt"

func panicTest() {
	panic("I am panicking!!")
}

func main() {
	panicTest()
	output := recover()

	fmt.Println(output)
}
