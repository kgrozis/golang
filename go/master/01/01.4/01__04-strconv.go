package main

import (
	"fmt"
	"os"      // cli
	"strconv" // convert cli string to numeric
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please give one or more floats.")
		os.Exit(1)
	}

	arguments := os.Args // slice, first element is name of program
	// assumes floats:
	min, _ := strconv.ParseFloat(arguments[1], 64) // access 2nd element of slice
	max, _ := strconv.ParseFloat(arguments[1], 64) // _ ignores/discards error value if not float

	for i := 2; i < len(arguments); i++ { // iterate thru cli args
		n, _ := strconv.ParseFloat(arguments[i], 64)
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	fmt.Println("Min:", min)
	fmt.Println("Max:", max)

}
