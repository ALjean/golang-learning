package main

import (
	"fmt"
	"time"
)

func timeWithTimeout(quit chan int, timer *time.Timer) {
	time.Sleep(1 * time.Second)
	select {
	case <-timer.C:
		fmt.Println("Time is over")

	case <-quit:
		if !timer.Stop() {
			<-timer.C
		}
		fmt.Println("Timer is stopped")
	default:
		fmt.Println("Function completion")
	}
}

func main() {
	timer := time.NewTimer(3 * time.Second)
	quit := make(chan int)

	go timeWithTimeout(quit, timer)

	quit <- 1

	time.Sleep(2 * time.Second)

}
