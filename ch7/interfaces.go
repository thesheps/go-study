package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

type triangle struct {
	a, b, c float64
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func measure(g geometry) {
	fmt.Println(g.area())
}

func main() {
	c := circle{10}
	r := rectangle{12, 13}
	t := triangle{1, 2, 3}
	measure(c)
	measure(r)

	//measure(t) This is naughty as triangle doesn't satisfy the interface.

}
