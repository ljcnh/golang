package main

import (
	"fmt"
	"time"
)

func main() {
	first := make(chan struct{})
	second := make(chan struct{})
	finall := make(chan struct{})
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		i := 0
	loop:
		for {
			second <- struct{}{}
			select {
			case <-first:
				i++
			case <-ticker.C:
				//3572262
				fmt.Printf("ping-pong: %d\n", 2*i/5)
				break loop
			}
		}
		finall <- struct{}{}
	}()
	go func() {
		for true {
			<-second
			first <- struct{}{}
		}
	}()
	<-finall
	fmt.Println()
}
