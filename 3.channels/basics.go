package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	for i := 1; i < 3; i++ {
		go func(i int) {

			for {
				time.Sleep(time.Second * 1)
				c <- i
				i++

			}

		}(i)
	}

	for i := range c {
		fmt.Println(i)
		// if i == 4 {
		// 	close(c)
		// }

	}

	time.Sleep(time.Second * 20)
}
