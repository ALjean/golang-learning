package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Client struct {
	Age int64
}

func addAge(client *Client, age int) {
	//mu.Lock()
	atomic.AddInt64(&client.Age, int64(age))
	//mu.Unlock()
}

func main() {
	client := &Client{}
	//mu := &sync.Mutex{}

	for i := 1; i <= 10000; i++ {
		go addAge(client, 1)
	}

	time.Sleep(1 * time.Second)
	fmt.Printf("%#v\n", client)
}
