package main

import "fmt"

func main() {
	// Initialises a dumb array:
	ages := make(map[string]int)
	ages["Steve"] = 1

	dobs := map[string]string{
		"Sheps": "1985-09-11",
	}

	fmt.Println(ages)
	fmt.Println(dobs)
}
