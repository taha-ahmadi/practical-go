package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go func() {
		for i := 0; i < 3; i++ {
			msg := fmt.Sprintf("message #%d", i+1)
			ch <- msg
		}
		close(ch)
	}()

	for msg := range ch {
		fmt.Println("got:", msg)
	}
	time.Sleep(10 * time.Second)
}
