package main

import (
	"fmt"
	"time"
//	"sync"
)

func main() {
//	var wg sync.WaitGroup
//	wg.Add(1)

//	go func() {
		// count("sheep")
		// wg.Done()
	// }()

	// wg.Wait()
	c := make(chan string)
	go count("sheep", c)

	for msg := range c {
		// IF we detect channel is closed, then break out FOR
		fmt.Println(msg)
	}
}

func count(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		// put thing in channel for reviever
		c <- thing
		// fmt.Println(i, thing)
		time.Sleep(time.Millisecond *500)
	}
	// Sender only can close the channel
	close(c)
}

