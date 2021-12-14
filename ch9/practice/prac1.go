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
	}()

	go func() {
		Deposit(25)
	}()

	go func() {
		Deposit(41)
	}()
	time.Sleep(5 * time.Second)
}

type withdrawMsg struct {
	amount  int
	success chan bool
}

var deposits = make(chan int)
var withdraw = make(chan withdrawMsg)
var balances = make(chan int)

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int {
	return <-balances
}

func Withdraw(amount int) bool {
	ch := make(chan bool)
	withdraw <- withdrawMsg{amount, ch}
	return <-ch
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case wd := <-withdraw:
			if balance >= wd.amount {
				balance -= wd.amount
				wd.success <- true
			} else {
				wd.success <- false
			}
		case balances <- balance:
			fmt.Println("balances")
		}
	}
}

func init() {
	go teller()
}
