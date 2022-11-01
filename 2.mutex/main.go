package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mutex sync.Mutex

type Income struct {
	Source string
	Amount int
}

func main() {

	// * variable for bank balance
	var bankBalance int

	// ! print out starting values
	fmt.Printf("Initial account balance: %d.00", bankBalance)
	fmt.Println()

	// ? define weakly revenue
	incomes := []Income{
		{Source: "Main job", Amount: 1100},
		{Source: "Part time job", Amount: 500},
		{Source: "Gifts", Amount: 742},
		{Source: "Investments", Amount: 1960},
	}

	// TODO: loop through 52 weeks (one year), and print out how much is made; keep running total

	wg.Add(len(incomes))

	for i, income := range incomes {
		go func(i int, income Income) {
			defer wg.Done()
			for week := 1; week <= 52; week++ {
				mutex.Lock()
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp
				fmt.Printf("On week %d, you earned $%d.00 from %s\n", week, income.Amount, income.Source)
				mutex.Unlock()
			}

		}(i, income)
	}

	wg.Wait()
	// print out final income

	fmt.Printf("Final bank balance: $%d.00", bankBalance)
	fmt.Println()

}
