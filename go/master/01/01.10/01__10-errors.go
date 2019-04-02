package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please give one or more floats.")
		os.Exit(1)
	}

	arguments := os.Args
	var err error = errors.New("an error") // create variable named error and initialize it
	k := 1
	var n float64

	for err != nil { // if first command-line arg is not a float, check the next arg
		if k >= len(arguments) {
			fmt.Println("None of the arguments is a float!") // no floats found
			return
		}
		n, err = strconv.ParseFloat(arguments[k], 64) // check if float
		k++

	}

	min, max := n, n // initialeze min, max values

	for i := 2; i < len(arguments); i++ {
		n, err := strconv.ParseFloat(arguments[i], 64)
		if err == nil {
			if n < min {
				min = n
			}
			if n > max {
				max = n
			}
		}
	}

	fmt.Println("Min:", min)
	fmt.Println("Max:", max)

}
