package main

import "fmt"

func readChan(ch chan int) {
	value := <-ch
	fmt.Println("Print chan ", value)
}

func main() {
	fmt.Println("Start main")
	var ch chan int

	ch = make(chan int)

	go readChan(ch)

	ch <- 77

	fmt.Println("End main")
}
