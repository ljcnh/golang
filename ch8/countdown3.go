package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	fmt.Println("Commencing countdown.  Press return to abort.")
	//tick := time.Tick(1 * time.Second)
	ticker := time.NewTicker(1 * time.Second)
	for countdowm := 10; countdowm > 0; countdowm-- {
		fmt.Println(countdowm)
		select {
		case <-abort:
			fmt.Println("launch aborted!")
			return
		case <-ticker.C:
		}
	}
	launch()
	ticker.Stop()
}
func launch() {
	fmt.Println("LIft off!")
}
