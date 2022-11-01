package main

import (
	"fmt"
	"sync"
)

var msg string

var wg sync.WaitGroup
var mutex sync.Mutex

func updateMessage(s string, m *sync.Mutex) {
	defer wg.Done()

	// lock msg, nothing can change value until we are done with it
	m.Lock()
	msg = s

	// unlock it
	m.Unlock()
}

func main() {
	// mutex = "mutual exclusion"
	// go run -race .
	// check if multiple go routines are trying to access same data (we should always do it when working with concurrency, to be safe)

	msg = "Hello W0rld!"

	wg.Add(2)

	go updateMessage("Hello, Africa!", &mutex)
	go updateMessage("Hello, Asia!", &mutex)

	wg.Wait()

	fmt.Println(msg)
}
