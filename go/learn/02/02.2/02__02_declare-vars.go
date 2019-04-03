// go is a strictly typed language.  Must bound a named element to both a value
//   and a type.
// Long form variable declaration:
//   var <named identifier> <type>
//   var can declare more than one variable identfiers of same type
// Variable declaraton and initialization:
//   var <named identifier> <type> = <value list or initializer expression>
//   var can initialize more than var identfier using commas
//   assignment is based on order of left and right side of =
// Omitting variable types:
//   var <named identifer> = <value list or initializer expression>
//   compiler infers variable type
// Short variable declaration:
//   <named identifer> := <value list or initializer expression>
//   loss var keyword and use assignment operator (:=)

package main

import "fmt"

// declare variable identfiers and type outside main
// if a value is not assigned go binds default value or zero-value
// package level scope for variables (global vars)
var name, desc string
var radius int32
var mass float64
var active bool = true  // initialize value as true
var satellites []string // List

// compiler infers var type:
var name2, desc2 = "Mars", "Planet" // infer type string
var radius2 = 6755                  // infer type int
var mass2 = 64169.0                 // infer type float

// Variable declaration block:
// increases readability
var (
	name4   = "Mars"
	radius4 = 6755
)

func main() {
	// print default values
	fmt.Println(name) // Value: "" (empty string)
	fmt.Println(desc)
	fmt.Println("Radius (km)", radius)    // Value: 0
	fmt.Println("Mass (kg", mass)         // Value: 0.0
	fmt.Println("Satellites", satellites) // Value: [ ] (Corresponds to array type)

	// assing value to declared variable
	// function block scope for variables
	name = "Sun"
	desc = "Star"
	radius = 685800
	mass = 1.989E+30
	active = true
	satellites = []string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}
	fmt.Println(name)
	fmt.Println(desc)
	fmt.Println("Radius (km)", radius)
	fmt.Println("Mass (kg)", mass)
	fmt.Println("Satellites", satellites)

	// short variable declaration:
	// must be assinged in a function block
	// must be used or compiler error
	// cannot update a previously decared variable
	// variable updates must be done with an equal sign (=)
	name3 := "Mars"
	radius3 := 6755
	mass3 := 64169.0
	fmt.Println(name3)
	fmt.Println("Radius (km)", radius3)
	fmt.Println("Mass (kg)", mass3)
}
