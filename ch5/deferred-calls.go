package main

import "fmt"

func main() {
	defer fmt.Println("I'll be called LAST, even though I'm FIRST!")
	fmt.Println("I'll be called FIRST, even though I'm LAST!")
}
