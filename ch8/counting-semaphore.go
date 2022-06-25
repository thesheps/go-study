package main

import (
	"fmt"
	"math/rand"
	"time"
)

var tokens = make(chan struct{}, 20)

func main() {
	for {
		go getRandomNumber()
	}
}

func getRandomNumber() int {
	tokens <- struct{}{}
	i := rand.Intn(5)

	time.Sleep(time.Duration(i) * time.Second)
	<-tokens

	fmt.Println(i)

	return i
}
