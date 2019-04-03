package main

import (
	"fmt"
)

// by default go supports UTF-8 for storing unicode values
// Rune type:
//   alias for the int32 type
//   intended to store unicode integer values encoded in utf-8
//   stored as a string literal constant surrounded by quotes
var (
	bksp = '\b'
	tab  = '\t'
	nwln = '\n'

	char1 = 'ɸ'
	char2 = 'আ'
	char3 = '語'

	char4 = '\u0369'
	char5 = '\xFA'
	char6 = '\045'
)

// string type:
//   a slice of immutable byte values
//   value of string is never changed
//   Unicode chars and encodes them using UTF-8
var (
	txt  = "水 and 火"
	txt2 = "\u6C34\x20brings\x20\x6c\x69\x66\x65."         // literal value enclosed in double quotes
	txt3 = `
	\u6C34\x20
	brings\x20
	\x6c\x69\x66\x65.
	` // literal value assigned by grave accents `.  prints raw string
)

func main() {
	// Rune converts chars to integers in output
	fmt.Println(bksp) // assigns integer value: 8
	fmt.Println(tab)  // assigns integer value: 9
	fmt.Println(nwln) // assigns integer value: 10

	fmt.Println(char1)
	fmt.Println(char2)
	fmt.Println(char3)

	fmt.Println(char4)
	fmt.Println(char5)
	fmt.Println(char6)

	fmt.Printf("%s (%d)\n", txt, len(txt)) // Prints the length of storage bytes used and not literal
	for i := 0; i < len(txt); i++ {
		fmt.Printf("%U", txt[i])
	}
	fmt.Println()
	fmt.Println(txt2)
	fmt.Println(txt3)

	txt = "This is a test"
	fmt.Println(txt)
}
