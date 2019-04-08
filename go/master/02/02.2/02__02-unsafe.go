package main

import (
	"fmt"
	"unsafe" // unsafe code bypasses the type safety and memory security of go
)

func main() {
	var value int64 = 5
	var p1 = &value
	var p2 = (*int32)(unsafe.Pointer(p1)) // unsasfe pointer from an int32 to int64

	fmt.Println("*p1: ", *p1) // astrick (*) dereferences a pointer and accesses value
	fmt.Println("*p2: ", *p2)
	*p1 = 5434123412312431212
	fmt.Println("value: ", value)
	fmt.Println("*p2: ", *p2) // output: *p2:  -930866580; 32b pointer cannot access a 64b value
	*p1 = 541234
	fmt.Println("value: ", value)
	fmt.Println("*p2: ", *p2)

}
