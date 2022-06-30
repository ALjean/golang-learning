package main

import (
	"errors"
	"fmt"
	"time"
)

func foo(timer *time.Timer, error chan<- error) {
	time.Sleep(1 * time.Second)
	select {
	case <-timer.C:
		fmt.Println("Timeout")
		error <- errors.New("timeout error")
	}

}

func main() {
	timer := time.NewTimer(2 * time.Second)
	error := make(chan error)

	go foo(timer, error)

	msg := <-error

	fmt.Printf("Error #%v", msg)

	time.Sleep(5 * time.Second)

}
