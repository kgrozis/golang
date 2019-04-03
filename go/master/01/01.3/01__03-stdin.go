package main

import (
	"bufio" // file io
	"fmt"
	"os" // os operations; platform independent
)

func main() {
	var f *os.File
	f = os.Stdin
	defer f.Close()

	scanner := bufio.NewScanner(f) // call newscanner using stdin
	for scanner.Scan() {           // read from stdin
		fmt.Println(">", scanner.Text()) // print each read line
	}

}
