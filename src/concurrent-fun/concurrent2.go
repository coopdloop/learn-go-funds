package concur2

import (
	"flag"
	"fmt"
	"os"
	"time"
)

var problem bool
var solution bool

func init() {
	// https://pkg.go.dev/flag
	flag.BoolVar(&problem, "problem", false, "set true, see problem case")
	flag.BoolVar(&solution, "solution", false, "set true, see solution case")
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	// Go routine #1
	go func() {
		for {
			// First Channel, send every 500 ms
			c1 <- "Every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()
	// Go routine #2
	go func() {
		for {
			// Second Channel, send every 2 s
			c2 <- "Every two seconds"
			time.Sleep(time.Second * 2)
		}
	}()

	flag.Parse()
	if problem && solution {
		fmt.Println("no args set")
		os.Exit(1)
	}

	// Check if user selected to see the problem output
	if problem {
		// PROBLEM CASE BELOW: FOR LOOP
		for {
			// inf loop, recieve from c1 and c2
			fmt.Println(<-c1)
			// c2 sends slower, forcing c1 to be blocked and wait, even though its ready for reciever.
			fmt.Println(<-c2)
		}
	}

	// Check if user selected to see the solution output
	if solution {
		// SOLVE CASE BELOW: FOR LOOP (select)
		for {
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			}
		}
	}
}
