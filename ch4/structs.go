package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Point
	Radius, Spokes int
}

// This simple example shows anonymous embedding of structs. Neato!
func main() {
	c := Circle{Point{1, 2}, 3}
	w := Wheel{Point{1, 2}, 3, 4}

	fmt.Println(c)
	fmt.Println(w)
}
