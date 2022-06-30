package main

import (
	"context"
	"fmt"
	"time"
)

func foo(ctx context.Context, num int) {
	select {
	case <-ctx.Done():
		fmt.Printf("Process #%v canceled\n", num)
	}

}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	for i := 0; i < 10; i++ {
		go foo(ctx, i)
	}

	time.Sleep(10 * time.Second)

}
