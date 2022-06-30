package main

import "fmt"

func writeChan(ch chan<- int) {
	for i := 0; i <= 5; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	fmt.Println("Start main")
	var ch chan int

	ch = make(chan int)

	go writeChan(ch)

	for i := range ch {
		fmt.Println("Print i: ", i)
	}

}
