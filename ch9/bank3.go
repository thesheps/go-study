package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	mu       sync.Mutex
	balance  int
	useMutex = false
)

func Deposit(amount int) {
	if useMutex {
		mu.Lock()
	}

	balance = balance + amount

	if useMutex {
		mu.Unlock()
	}
}

func Balance() {
	if useMutex {
		mu.Lock()
	}

	b := balance

	if useMutex {
		mu.Unlock()
	}

	fmt.Println(b)
}

func main() {
	go func() { Deposit(10) }()
	go func() { Deposit(10) }()
	go func() { Deposit(10) }()
	go func() { Deposit(10) }()
	go func() { Deposit(10) }()
	go func() { Deposit(10) }()
	go func() { Deposit(10) }()
	go func() { Deposit(10) }()
	go func() { Deposit(10) }()
	go func() { Deposit(10) }()
	go func() { Deposit(10) }()
	go func() { Deposit(10) }()
	go func() { Deposit(10) }()
	go func() { Deposit(10) }()
	go func() { Deposit(10) }()
	go func() { Deposit(10) }()
	go func() { Deposit(10) }()
	go func() { Deposit(10) }()
	go func() { Deposit(10) }()
	go func() { Deposit(10) }()
	go func() { Deposit(10) }()
	go func() { Deposit(10) }()
	go func() { Deposit(10) }()
	go func() { Deposit(10) }()
	go func() { Deposit(10) }()
	go func() { Deposit(10) }()
	go func() { Deposit(10) }()
	go func() { Deposit(1) }()

	time.Sleep(1 * time.Second)
	Balance()
}
