package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	count := 0

	for tick := range ticker.C {
		count++
		fmt.Printf("Tick #%v, time %v\n", count, tick)
		if count > 5 {
			ticker.Stop()
			break
		}
	}
}
