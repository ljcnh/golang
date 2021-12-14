package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		Deposit(200)
		fmt.Println("=", Balance())
	}()

	go func() {
		Deposit(100)
		fmt.Println("=", Balance())
	}()

	go func() {
		Deposit(25)
		fmt.Println("=", Balance())
	}()

	go func() {
		Deposit(41)
		fmt.Println("=", Balance())
	}()
	time.Sleep(5 * time.Second)
	fmt.Println("=", Balance())
}

var deposits = make(chan int)
var balances = make(chan int)

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int {
	return <-balances
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func init() {
	go teller()
}
