package main

import "fmt"

func main() {
	x := 1
	y := 2

	x, y = y, x

	fmt.Println(x, y)
}
