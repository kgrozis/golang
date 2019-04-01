package main

import (
	"fmt"
	"math"
)

// all parameters passed to a func are done so by value
//   a local copy of the passed value is created in the func
//   do not pass by reference
func dbl(val float64) { // val is local copy of p
	val = 2 * val // update param
	fmt.Printf("dbl()=%.5f\n", val)
}

// modify the value referenced by pointer
func half(val *float64) { // create local copy of ptr value
	fmt.Printf("call half(%f)\n", *val) // ptr operator dereferences and manipulates value
	*val = *val / 2                     // changes will be accessible in main()
}

// anonymous functions and closures
//   funcs declared and bound to vars (mul, sqr)
//   return value becomes var
var (
	mul = func(op0, op1 int) int {
		return op0 * op1
	}
	sqr = func(val int) int {
		return mul(val, val)
	}
)

func main() {
	p := math.Pi
	fmt.Printf("before dbl()=%.5f\n", p) // 3.14...
	dbl(p)
	fmt.Printf("after dbl()=%.5f\n", p) // p is unchanged

	// pass-by-reference
	//   use pointer parameter values
	//   func looks outside its scope and changes the value stored at the location
	//   referenced by the pointer
	num := 2.807770
	fmt.Printf("\nnum=%f\n", num)
	half(&num) // send ptr to num value
	fmt.Printf("half(num)=%f\n", num)

	// anonymous functions and closures
	fmt.Printf("\nmul(24,7) = %d\n", mul(25, 7))
	fmt.Printf("sqr(13 = %d\n", sqr(13))

	// invoking anonymous func literals
	//   func evaluated in place as an exprssion that returns the func result
	fmt.Printf(
		"\n94 (°F) = %.2f (°C)\n",
		func(f float64) float64 { // func invoked with Printf, normal func signature
			return (f - 32.0) * (5.0 / 9.0)
		}(94), // func args (var f)
	)

	// closures
	//   func literals have lexical visibility ot non-local vars declared outside
	//   of there enclousing code block
	fmt.Printf("\n")
	for i := 0.0; i < 360.0; i += 45.0 {
		// func literal code block
		//   each iteration of loop, a closure is formed between the enclosed
		//   literal and outer non-local var i
		//   can access non-local values without other means like ptrs
		rad := func() float64 {
			return i * math.Pi / 180
		}()
		fmt.Printf("%.2f Deg = %.2f Rad\n", i, rad)
	}
}
