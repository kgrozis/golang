package main

import "fmt"

var (
	int1, int2, int3 int     = 1, 2, 3
	float1, float2   float32 = 3.5, 5.25
	string1, string2 string  = "Go", "Programming"
)

// Go operators do not support operator-overloading
func main() {
	// Common Arithmetic operators
	//   *,/,- support ints, floats, and complex numbers
	fmt.Println(int1 * int2)
	fmt.Println(int2 / int1)
	fmt.Println(int2 - int1)
	fmt.Println(float1 * float2)
	fmt.Println(float2 / float1)
	fmt.Println(float2 - float1)

	// Remainder supports ints
	fmt.Println(int2 % int1)

	// Addition supports ints, floats, and strings
	fmt.Println(int1 + int2)
	fmt.Println(float1 + float2)
	fmt.Println(string1 + " " + string2)

	// increment (++) and decrement (--) operators work like C
	//   ++ and -- are statements and not expressions
	//   must be applied to suffix and will not compile when appended to var
	for i := len(string1) - 1; i >= 0; {
		fmt.Println(string(string1[i])) // Prints inverse string
		i--
	}

	// Assignment Operators
	// applies the operation on var and saves with var
	int1 += 10
	int2 -= 10
	float1 *= 10
	float2 *= 10
	int3 %= 1
	fmt.Println(int1, int2, float1, float2, int3)

}
