package main

import (
	"flag"
	"fmt"
	"learn-go-funds/src/concurrent0"
	"learn-go-funds/src/concurrent1"
	"learn-go-funds/src/concurrent2"
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
		fmt.Println(concur0.main)
	}

}
