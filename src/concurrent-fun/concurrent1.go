package concur1

import (
	"fmt"
)

func main() {
	c := make(chan string, 2)
	c <- "hello"
	c <- "world"
	// Creates Deadlock
	// c <- "third"

	msg := <-c
	fmt.Println(msg)
	msg = <-c
	fmt.Println(msg)
}