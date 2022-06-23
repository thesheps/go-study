package main

import "fmt"

func main() {
	if x := f(); x == 0 {
		fmt.Println(x)
	} else if x := f(); x == 0 {
		fmt.Println(x)
	}

	fmt.Println(x)
}

func f() int {
	return 1
}
