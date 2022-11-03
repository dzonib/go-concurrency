package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const NumberOfPizzas = 10

var pizzasMade, pizzasFailed, total int

// ? Channels are far more powerful then wait groups and mutexes, because
// ? They allow to exchange data between the go routines

// Producer is a type for structs that holds two channels: one for pizzas, with all
// information for a given pizza order including whether it was made
// successfully, and another to handle end of processing (when we quit the channel)
type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

// PizzaOrder is a type for structs that describes a given pizza order. It has the order
// number, a message indicating what happened to the order, and a boolean
// indicating if the order was successfully completed.
type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

// Close is simply a method of closing the channel when we are done with it (i.e.
// something is pushed to the quit channel)
func (p *Producer) Close() error {
	ch := make(chan error)

	p.quit <- ch
	return <-ch
}

// makePizza attempts to make a pizza. We generate a random number from 1-12,
// and put in two cases where we can't make the pizza in time. Otherwise,
// we make the pizza without issue. To make things interesting, each pizza
// will take a different length of time to produce (some pizzas are harder than others).
func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++

	if pizzaNumber <= NumberOfPizzas {
		delay := rand.Intn(5) + 1
		fmt.Printf("Received order number %d!\n", pizzaNumber)

		rnd := rand.Intn(12) + 1

		msg := ""
		success := false

		if rnd < 5 {
			pizzasFailed++
		} else {
			pizzasMade++
		}

		total++
		fmt.Printf("Making pizzas #%d. Will take %d seconds...\n", pizzaNumber, delay)

		// delay for a bit
		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf("*** We ran out of ingredients for pizza number %d!", pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("*** The cook quit, while making pizza number %d!", pizzaNumber)
		} else {
			success = true
			msg = fmt.Sprintf("Pizza order number #%d is ready! :)", pizzaNumber)
		}

		p := PizzaOrder{
			pizzaNumber: pizzaNumber,
			message:     msg,
			success:     success,
		}

		return &p
	}

	// this will be 11... (we only should have 10)
	return &PizzaOrder{pizzaNumber: pizzaNumber}
}

// pizzeria is a goroutine that runs in the background and
// calls makePizza to try to make one order each time it iterates through
// the for loop. It executes until it receives something on the quit
// channel. The quit channel does not receive anything until the consumer
// sends it (when the number of orders is greater than or equal to the
// constant NumberOfPizzas).
func pizzeria(pizzaMaker *Producer) {
	// keep track of which pizza we are making
	var i = 0

	// this loop will continue to execute, trying to make pizzas,
	// until the quit channel receives something.
	for {
		currentPizza := makePizza(i)
		if currentPizza != nil {
			i = currentPizza.pizzaNumber

			// select are used only for channels
			// they function similar as switch statement

			select {
			// we tried to make a pizza (we sent something to the data channel - a chan PizzaOrder)
			case pizzaMaker.data <- *currentPizza:

			case quitChan := <-pizzaMaker.quit:

				// close channels
				close(pizzaMaker.data)
				close(quitChan)

				// exit go routine
				return
			}
		}
	}
}

func main() {
	//? Seed the random number generator
	rand.Seed(time.Now().Unix())

	//! Print out the message
	color.Cyan("The Pizzeria is open for business")
	color.Cyan("---------------------------------")

	//? Create producer
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	//TODO run the producer as go routine in background
	go pizzeria(pizzaJob)

	//? Create and run consumer (customer)

	for i := range pizzaJob.data {
		if i.pizzaNumber <= NumberOfPizzas {

			if i.success {
				color.Green(i.message)
				color.Green("Order #%d is out for delivery", i.pizzaNumber)
			} else {
				// we tried to make pizza and something went wrong
				color.Red(i.message)
				color.Red("The customer is really mad!")
			}
		} else {
			color.Cyan("Done making pizzas...zZzZz...")
			err := pizzaJob.Close()

			if err != nil {
				color.Red("*** Error closing channel", err)
			}
		}
	}

	//! Print out the ending message
	color.Cyan("--------------")
	color.Cyan("Done for a day")
	color.Cyan("We made a %d pizzas, but failed to make %d  with %d attempts in total", pizzasMade, pizzasFailed, total)

	switch {
	case pizzasFailed > 9:
		color.Red("It was an awful day :(")
	case pizzasFailed >= 6:
		color.Red("It was not a weary good day :/")
	case pizzasFailed >= 4:
		color.Yellow("It was an OK day!")
	case pizzasFailed >= 2:
		color.Green("It was a weary good day! :)")
	default:
		color.Green("It was a great day !! :D")
	}
	color.Cyan("--------------")
}
