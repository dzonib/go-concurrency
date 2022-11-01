package main

import (
	"fmt"
	"sync"
)

var msg string

func updateMessage(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	msg = s
	printMessage()
}

func printMessage() {
	fmt.Println(msg)
}

var wg sync.WaitGroup

func main() {

	// challenge: modify this code so that the calls to updateMessage() on lines
	// 28, 30, and 33 run as goroutines, and implement wait groups so that
	// the program runs properly, and prints out three different messages.
	// Then, write a test for all three functions in this program: updateMessage(),
	// printMessage(), and main().

	greetings := []string{
		"Hello, universe!",
		"Hello, cosmos!",
		"Hello, world!",
	}

	wg.Add(len(greetings))

	for i, g := range greetings {
		go updateMessage(fmt.Sprintf("%d: %s", i, g), &wg)

	}

	wg.Wait()
}
