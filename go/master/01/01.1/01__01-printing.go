package main

import "fmt"

func main() {
	v1 := "123"
	v2 := 123
	v3 := "Have a nice day\n" // new line char
	v4 := "abc"               // no new line

	fmt.Print(v1, v2, v3, v4)                      // no spacing
	fmt.Println()                                  // prints new line char
	fmt.Println(v1, v2, v3, v4)                    // prints spacing and newline
	fmt.Print(v1, " ", v2, " ", v3, " ", v4, "\n") // user adds spacing and newline
	fmt.Printf("%s%d %s %s\n", v1, v2, v3, v4)     // with formatting

}
