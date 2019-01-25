package main

import "fmt"

// defer keyword:
//   using defer keyword before func keyword
//   delays execution until before enclosing function returns

// using defer
//   resource clean-up code:
//   - close files
//   - release net resources
//   - close go channel
//   - commit db transactions
//   etc...

func do(steps ...string) { // steps: variadic parameters
	defer fmt.Println("All done!")

	for _, s := range steps {
		// last deferred statement is executed first (start car...)
		defer fmt.Println(s) // defer each element
	}

	fmt.Println("Starting")
}

func main() {
	do(
		"Find key",
		"Apply break",
		"Put key in ignition",
		"Start car",
	)
}
