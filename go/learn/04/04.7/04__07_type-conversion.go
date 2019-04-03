package main

import "fmt"

// go considers each type to be different
//   build errors will occur on type mismatch
//   this includes of similar types: int, int32, int64 are mismatched
//   compiler will reject declared types with same underlying types
// type conversion expression:
//   <target_type>(<value or expression>)

func main() {
	var count int32
	var actual int
	// conversion of int to int32 -- actual, match expression
	var test int32 = int32(actual) + count
	fmt.Println(test) // nil

	// target type and converted value are both simple numeric types
	var i int
	var i2 int32 = int32(i)
	var re float64 = float64(i + int(i2))
	fmt.Println(re) // nil

	// target type and converted value are both complex numeric types
	var cn64 complex64
	var cn128 complex128 = complex128(cn64)
	fmt.Println(cn128) // nil

	// target type and converted value have the same underlying type
	type signal int
	var sig signal
	var event int = int(sig)
	fmt.Println(event) // nil

	// target type is a string and the converted value is a valid integer type
	a := string(72)
	b := string(int32(101))
	c := string(rune(108))
	fmt.Println(a, b, c) // H e l

	// target type is string and the coverted value is a slice of bytes, int32, or runes
	msg0 := string([]byte{'H', 'i'})
	msg1 := string([]rune{'Y', 'o', 'u', '!'})
	fmt.Println(msg0, msg1) // Hi You!

	// target type is a slice of byte, int32, or rune values and converted value is a string
	data0 := []byte("Hello")
	data1 := []int32("World!") // [72 101 108 108 111] [87 111 114 108 100 33]
	fmt.Println(data0, data1)
}
