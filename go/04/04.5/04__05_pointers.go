package main

import (
	"fmt"
	"strings"
)

// pointer
//   used to indirectly reference memory addresses where data are stored
//   more efficent way of accessing data value
//   no pointer arithmetic
// the pointer type:
//   Uses * operator to designate a type as a pointer
//   a variable of type T uses expression *T as its pointer type
//   can only be assigned values of their declared type
//   if nothing is referenced, it is represented by the literal contant nil
var valPtr *float32  // nil
var countPtr *int    // nil
var person *struct { // nil
	name string
	age  int
}
var matrix *[1024]int // nil
var row []*int64      // nil

func double(x *int) {
	*x = *x * 2 // dereference *x value ot 3 and update *x to 6
}

func cap(p *struct{ first, last string }) {
	p.first = strings.ToUpper(p.first) // dereference *p to efg and update first
	p.last = strings.ToUpper(p.last)   // dereference *p ot efg and update last
}

func main() {
	fmt.Println(valPtr, countPtr, person, matrix, row)

	// address operator (&) obtains address value of a variable
	var a int = 1024
	var aptr *int = &a // assign aptr to address of a
	fmt.Printf("a=%v\n", a)
	fmt.Printf("aptr=%v\n", aptr)

	structPtr := &struct{ x, y int }{44, 55}
	pairPtr := &[2]string{"A", "B"}
	fmt.Printf("struct=%#v, type=%T\n", structPtr, structPtr)
	fmt.Printf("pairPtr=%#v, type=%T\n", pairPtr, pairPtr)

	// new() function
	//   can be used intialize a pointer value
	//   allocs memoryfor a zero-value of specified type
	//   return address for newly created value
	intPtr := new(int)                     // initialize to nil
	*intPtr = 44                           // update value to 44
	p := new(struct{ first, last string }) // initialize to nil
	p.first = "Samuel"                     // update struct value
	p.last = "Pierre"                      // update struct value
	fmt.Printf("Value %d, type%T\n", *intPtr, intPtr)
	fmt.Printf("Person %+v\n", p)

	// Pointer indirection
	//   If you have an address, you can access value with * operator
	//   to point value
	abc := 3
	double(&abc)
	fmt.Println(abc)
	efg := &struct{ first, last string }{"Max", "Planck"}
	cap(efg)
	fmt.Println(efg)
}
