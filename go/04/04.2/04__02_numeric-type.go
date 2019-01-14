package main

import (
	"fmt"
	"math"
	"unsafe"
)

// signed integer types
var _ int8 = 12       // signed 8b number, -128-127
var _ int16 = -400    // signed 16b number, -32768-32767
var _ int32 = 12022   // signed 32b number, -2147483648-214748367
var _ int64 = 1 << 33 // signed 64b number,-9223372036854775808 - 9223372036854775807
var _ int = 3 + 1415  // implementation specific either 32b or 64b signed int

// unsigned integer types
var _ uint8 = 18                 // unsigned 8b number, 0-255
var _ uint16 = 44                // unsigned 16b number, 0-65535
var _ uint32 = 133121            // unsigned 32b number, 0-4294967295
var _ uint64 = 23113233          // unsigned 64b number, 0-18446744073709551615
var _ uint = 7542                // implementation specific either 32b or 64b unsigned int
var _ byte = 255                 // unsigned 8b number
var i uint64 = 123456            // var for pointer
var _ uintptr = unsafe.Sizeof(i) // unsigned, designed to store pointers

// floating point types
var _ float32 = 0.5772156649 // signed 32b, ieee-754 single precision floating point values
var _ float64 = math.Pi      // signed 64b, ieee-754 double precision floating point values

// complex number types
var _ complex64 = 3.5 + 2i // float32, real and imaginary
var _ complex128 = -5.0i   // float64, real and imaginary

func main() {
	fmt.Println("all types declared!")

	// numberic literals
	//   hexadecimals are represented as 0x or 0X
	//   octal values start with 0
	vals := []int{
		1024,       // decimal
		0x0FF1CE,   // hex
		0x8BADF00D, // hex
		0xBEEF,     // hex
		0777,       // oct
	}
	for _, i := range vals {
		if i == 0xBEEF {
			fmt.Printf("Got %d\n", i)
			break
		}
	}

	// float exponenetial notation
	//   can be represented with e or E
	//   e or E translated into <float> * 10^<#>
	p := 3.1415926535 // float decimal notation
	e := .5772156649
	x := 7.2E-5       // float exponential notation
	y := 1.616199e-35 // float exponential notation
	z := .416833e32
	fmt.Println(p, e, x, y, z)

	// complex number notation
	//   imaginary literal is a float followed by the letter i
	//   builtin real() returns real portion of complex
	//   builtin imag() returns imaginary portion of complex
	a := -3.5 + 2i
	fmt.Printf("%v\n", a)
	fmt.Printf("%+g, %+g\n", real(a), imag(a))

}
