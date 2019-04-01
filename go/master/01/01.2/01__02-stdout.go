package main

import (
	"io" // use io instead of fmt for outputing
	"os" // use to read cli arguments
)

func main() {
	myString := "" // holds text to print
	arguments := os.Args

	if len(arguments) == 1 {
		myString = "Please give me one argument!"
	} else {
		myString = arguments[1] // first cli argument
	}

	io.WriteString(os.Stdout, myString) // works the same as fmt.Print()
	io.WriteString(os.Stdout, "\n")

}
