package main

import (
	"fmt"
	"runtime"
	"time"
)

func WorkAndPrint(num int) {
	fmt.Printf("START JOB #%v\n", num)
	var calc int
	for i := 0; i < 1000; i++ {
		calc = i * num
	}
	runtime.Gosched()
	fmt.Printf("FINISH JOB #%v: calc #%v\n", num, calc)
}

func main() {

	for i := 1; i <= 5; i++ {
		go WorkAndPrint(i)
	}

	// Cooperative multithreading
	// Goroutine does not require a context switch on the OS side, because it is so lightweight
	// The GO scheduler manages the goroutines by itself
	// scheduler manager
	// P - Processor (as many threads as available cores on your computer)
	// M (machine) - Machine Tread OS
	// G
	time.Sleep(100 * time.Millisecond)
}
