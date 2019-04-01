package main

import "fmt"

type Currency struct {
	Name    string
	Country string
	Number  int
}

var CAD = Currency{
	Name:    "Canadian Dollar",
	Country: "Canada",
	Number:  124,
}

var FJD = Currency{
	Name:    "Fiji Dollar",
	Country: "Fiji",
	Number:  242,
}

var JMD = Currency{
	Name:    "Jamaicon Dollar",
	Country: "Jamaica",
	Number:  388,
}

var USD = Currency{
	Name:    "US Dollar",
	Country: "USA",
	Number:  840,
}

func main() {
	num0 := 242
	// parentheses around test expression are not needed
	// Curly braces are required, even if a single statement
	// best practices is to write an if statement on multiple lines
	//    regardless of the simplicity
	if num0 > 100 || num0 < 900 {
		fmt.Println("Currency:", num0)
		printCurr(num0)
	} else {
		fmt.Println("Currency unknown")
	}

	// if statement is proceeded by an initialization statement
	// initialization statment is executed before test expression
	if num1 := 388; num1 > 100 || num1 < 900 {
		fmt.Println("Currency:", num1)
		printCurr(num1)
	}
}

func printCurr(number int) {
	if CAD.Number == number {
		fmt.Printf("Found: %+v\n", CAD)
		// if else if chain
		// switch statements are cleaner versions of if else if chains
	} else if FJD.Number == number {
		fmt.Printf("Found: %+v\n", FJD)
	} else if JMD.Number == number {
		fmt.Printf("Found: %+v\n", JMD)
	} else if USD.Number == number {
		fmt.Printf("Found: %+v\n", USD)
		// else statements are optional
		// only evaluated with the if statement(s) is false
	} else {
		fmt.Println("No currency found with number", number)
	}
}
