package main

import (
	"fmt"
	"strings"
)

type Curr struct {
	Currency string
	Name     string
	Country  string
	Number   int
}

var currencies = []Curr{
	Curr{"DZD", "Algerian Dinar", "Algeria", 12},
	Curr{"AUD", "Australian Dollar", "Australia", 36},
	Curr{"EUR", "Euro", "Belgium", 978},
	Curr{"CLP", "Chilean Peso", "Chile", 152},
	Curr{"EUR", "Euro", "Greece", 978},
	Curr{"HTG", "Gourde", "Haiti", 332},
}

func isDollar(curr Curr) bool {
	var result bool
	switch curr {
	// switch statement creates multi-way branching by evaluating values or
	//   expressions from case clauses
	// 'break' statement can be used to escape out of switch code block early
	// expression switch: test values of any types
	//   this switch block is testing var curr, type Curr
	//   evaluate left to right, top to bottom until a match is found
	//   execute statements for matching expression and exit switch code block
	default:
		result = false
	case Curr{"AUD", "Australian Dollar", "Australia", 36}:
		result = true
	case Curr{"HKD", "Hong Kong Dollar", "Hong Kong", 344}:
		result = true
	case Curr{"USD", "US Dollar", "United States", 840}:
		result = true
	}
	return result
}
func isDollar2(curr Curr) bool {
	dollars := []Curr{currencies[0], currencies[1], currencies[2]}
	switch curr {
	default:
		return false
	case dollars[0]:
		fallthrough
		// 'fallthrough' causes Go to move to the next executable statement
	case dollars[1]:
		fallthrough
		// 'fallthrough' causes Go to move to the next executable statement
	case dollars[2]:
		return true
		// dollars[0-2] will return 'true' because of the fallthrough statement
	}
	return false
}

func isEuro(curr Curr) bool {
	switch curr {
	case currencies[2], currencies[4], currencies[5]:
		// commas act as logical OR operators
		// reduces complexity in a similar If statement
		return true
	default:
		return false
	}
}

func find(name string) {
	for i := 0; i < 6; i++ {
		c := currencies[i]
		// expressionless switch
		//   acts similar to if statement with multiple OR operators
		switch {
		case strings.Contains(c.Currency, name),
			strings.Contains(c.Name, name),
			strings.Contains(c.Country, name):
			fmt.Println("Found", c)
		}
	}
}

func findNumber(num int) {
	for _, curr := range currencies {
		if curr.Number == num {
			fmt.Println("Found", curr)
		}
	}
}

func findAny(val interface{}) {
	// type switch: queries based on type information
	//   compare underlying value (or expressions) type information
	switch i := val.(type) {
	case int:
		findNumber(i)
	case string:
		find(i)
	default:
		fmt.Printf("Unable to search with type %T\n", val)
	}
}

func assertEuro(c Curr) bool {
	// switch keyword is followed by an initializer for vars local to switch code
	//   block are declared and initialized
	//   below initialize name and curr
	// expressionless switch statement
	switch name, curr := "Euro", "EUR"; {
	case c.Name == name:
		return true
	case c.Currency == curr:
		return true
	}
	return false
}

func main() {
	curr := Curr{"AUD", "Australian Dollar", "Australia", 36}
	if isDollar(curr) {
		fmt.Printf("%+v is Dollar currency\n", curr)
	} else if isEuro(curr) {
		fmt.Printf("%+v is Euro currency\n", curr)
	} else {
		fmt.Println("Currency is not Dollar or Euro")
	}
	dol := Curr{"HKD", "Hong Kong Dollar", "Hong Koong", 344}
	if isDollar2(dol) {
		fmt.Println("Dollar currency found:", dol)
	}

	var pass_to_find string = "AUD"
	find(pass_to_find)

	curr1 := Curr{"EUR", "Euro", "Italy", 978}
	if assertEuro(curr1) {
		fmt.Println(curr1, "is Euro")
	}

	findAny("Peso")
	findAny(36)
	findAny(false)
}
