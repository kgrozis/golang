package main

import "fmt"

// go methods:
//   a method is defining a func with a scope of a type (or attached to a type)
//   defined like any other func in go, however def includes a method receiver
//   method reciever is placed before method's name, to specify type attached
//   methods can only be accessed via declared type value (concrete or pointer)

// gallon base type:
//   cannot be a pointer or interface
type gallon float64

// method def:
//   reciver: (g gallon)
//   access method: gal.quart()
//   value g recieves the gal value 5 as local value g
// value and pointer receivers
//   receivers are normal func parameters thus pass-by-value applies and method get a full copy of original value
//   can pass parameters as values or pointers of base type
func (g gallon) quart() quart {
	return quart(g * 4)
}
func (g gallon) half() {
	g = gallon(g * 0.5) // update copy of gal in local value g; no return
}
func (g *gallon) double() {
	*g = gallon(*g * 2) // pointer refers to declared value and updates gal
}

// ounce base type:
type ounce float64

func (o ounce) cup() cup {
	return cup(o * 0.1250)
}

// cpu base type:
type cup float64

func (c cup) quart() quart {
	return quart(c * 0.25)
}
func (c cup) ounce() ounce {
	return ounce(c * 8.0)
}

// quart base type:
type quart float64

func (q quart) gallon() gallon {
	return gallon(q * 0.25)
}
func (q quart) cup() cup {
	return cup(q * 4.0)
}

func main() {
	gal := gallon(5)
	fmt.Printf("%.2f gallons = %.2f quarts\n", gal, gal.quart())

	// convert gallons to ounces by chaining methods
	// return type important to next link in chain
	ozs := gal.quart().cup().ounce()
	fmt.Printf("%.2f gallons = %.2f ounces\n", gal, ozs)

	// updates copy stored as g but not declared value of gal
	gal.half()
	fmt.Println(gal) // prints 5
	// pointer refers to declared value and updates gal
	gal.double()
	fmt.Println(gal) // prints 10
}
