package main

import "fmt"

func main() {
	// boolean type
	// boolean binary values are stored in bool
	// 1B value
	// 2 pre-declared literals: true and false represent bool values
	var readyToGo bool = false
	if !readyToGo {
		fmt.Println("Come on")
	} else {
		fmt.Println("Let's go!")
	}
}
