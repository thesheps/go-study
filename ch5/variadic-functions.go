package main

import "fmt"

func sums(vals ...int) int {
	var total int

	for _, val := range vals {
		total += val
	}

	return total
}

func main() {
	output := sums(1, 56, 3, 236, 347, 235, 3467, 3)

	fmt.Println(output)
}
