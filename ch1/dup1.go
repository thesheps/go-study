package main

import (
	"bufio"
	"fmt"
	"os"
)

// This example doesn't work by default, given that input.Scan() always returns true.
// Instead we just check to see if the length of the string is 0 and quit early
func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		test := input.Text()

		if len(test) == 0 {
			break
		}

		counts[test]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s", n, line)
		}
	}
}
