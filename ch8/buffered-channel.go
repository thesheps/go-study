package main

import (
	"fmt"
	"time"
)

var responses = make(chan string, 3)

func main() {
	response := mirroredQuery()
	fmt.Println(response)
}

func mirroredQuery() string {
	go request1()
	go request2()
	go request3()

	return <-responses
}

func request1() {
	time.Sleep(100 * time.Second)
	responses <- "Foo"
}

func request2() {
	time.Sleep(5 * time.Second)
	responses <- "Bar"
}

func request3() {
	time.Sleep(10 * time.Second)
	responses <- "Baz"
}
