package main

import (
	"fmt"
	"os"
)

// os.args
//   makes all command-line args available
//   creates a slice with of args using space to deliminate
//   position 0 holds the fully qualified name of programs binary path

func main() {
	for _, arg := range os.Args {
		fmt.Println(arg)
	}
}
