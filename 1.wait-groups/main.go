package main

import (
	"fmt"
	"sync"
)

// it is much safer to use pointers to wait group
func printSomething(s string, wg *sync.WaitGroup) {
	// decrements wait group by one
	defer wg.Done()

	fmt.Println(s)
}

func main() {
	// will not run, programs finished to fast for it to be executed
	// opens new lightweight thread and runs the function there
	// go printSomething("This is the first thing to be printed")

	var wg sync.WaitGroup

	words := []string{
		"alpha",
		"beta",
		"delta",
		"gama",
		"Epsilon",
		"Zeta",
		"Eta",
		"Theta",
		"Last",
	}

	wg.Add(len(words))

	for i, l := range words {
		go printSomething(fmt.Sprintf("%d: %s", i, l), &wg)
	}

	wg.Wait()

	// time.Sleep(1 * time.Second)

	wg.Add(1)
	printSomething("This is the second thing to be printed", &wg)
}
