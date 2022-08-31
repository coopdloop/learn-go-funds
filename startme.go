package main

import (
	"flag"
	"fmt"
	"github.com/coopdlop/learn-go-funds"
)

var concur0 bool

func init() {
	// https://pkg.go.dev/flag
	flag.BoolVar(&concur0, "concur0", false, "set true, see solution case")
	//var concur1 := flag.Bool("concur1", false, "set true, see solution case")
	//var concur2 := flag.Bool("concur2", false, "set true, see solution case")
}

func main() {

	flag.Parse()

	// Check if user selected to see the problem output
	if concur0 {
		fmt.Println(concur1.main)
	}

}
