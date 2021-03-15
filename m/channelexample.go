package main

import (
	"fmt"
	"time"
)

func main() {
	events := make(chan int)

	go func() {
		for i := 1; i <= 10; i++ {
			events <- i
			time.Sleep(250 * time.Millisecond)
		}
		close(events)
	}()

	for e := range events {
		go func() {
			fmt.Println(e)
		}()
	}
}
