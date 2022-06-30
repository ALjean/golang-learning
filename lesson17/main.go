package main

import (
	"context"
	"fmt"
	"time"
)

func sendData(ctx context.Context, num int) {
	timer := time.NewTimer((time.Duration(num)) * time.Second)

	select {
	case <-ctx.Done():
		fmt.Printf("Process #%v canceled\n", num)
	case <-timer.C:
		fmt.Printf("Data #%v send\n", num)
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	for i := 1; i <= 10; i++ {
		go sendData(ctx, i)
	}
	time.Sleep(1 * time.Second)
	cancel()
	time.Sleep(500 * time.Millisecond)
}
