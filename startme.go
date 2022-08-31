package main

import (
	"flag"
	//"fmt"
	"concurrent-fun/pkg1"
)

var concurrent0 bool
var concurrent1 bool
var concurrent2 bool
func init() {
	// https://pkg.go.dev/flag
	flag.BoolVar(&concurrent0, "concurrent0", false, "set true, see solution case")
	flag.BoolVar(&concurrent1, "concurrent1", false, "set true, see solution case")
	flag.BoolVar(&concurrent2, "concurrent2", false, "set true, see solution case")

}

func main() {
	//pkg1.Concur0()
	flag.Parse()

	// Check if user selected to see the problem output
	if concurrent0 {
		pkg1.Concur0()
	}
	// Check if user selected to see the problem output
	if concurrent1 {
		pkg1.Concur1()
	}
			// Check if user selected to see the problem output
	if concurrent2 {
		pkg1.Concur2()
	}

}
