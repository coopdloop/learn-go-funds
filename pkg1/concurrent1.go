package pkg1

import (
	"fmt"
)

func Concur1() {
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
