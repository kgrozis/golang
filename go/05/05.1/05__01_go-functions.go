package main

import (
	"fmt"
	"math"
)

// functions:
//   first-class and typed elements
// function declaration:
//   func [<func-identifier>] ([<argument-list>] [(<result-list)]){
//     ...
//     [return] [<value or expression list>]
//   }
//   func - keyword marks the beginning of a func declaration
// func signature:
//   func identifier (optional) - name the func for future reference
//   argument-list - required set of parenthese enclosing an optional list of
//     func arguments
//   result-list (optional) - list of types for values returned by func.  when
//     func returns more than one value enclosing set of parentheses is required
//   return - causes the execution flow to exit a func. when the func defines
//     returned values in its def, a return statement is required as the last
//     logical statement of the func
//   identifies func.  can be different func with differing order of elements,
//      differing number of elements, differing returns.

// name identfier print Pi
// takes no parameter and returns no values (no return statement required)
func printPi() {
	fmt.Printf("printPi() %v\n", math.Pi)
}

// name identifier avogadro
// takes no parameter
// returns a value of type float64 (return statement required)
// return value is declared as part of the func signature
func avogadro() float64 {
	return 6.02214129e23
}

// name identifier fib
// takes parameter n of type int and prints the fibonacci seq for up to n
// takes no parameter and returns no values (no return statement required)
func fib(n int) {
	fmt.Printf("fib(%d): [", n)

	var p0, p1 uint64 = 0, 1
	fmt.Printf("%d %d ", p0, p1)
	for i := 2; i <= n; i++ {
		p0, p1 = p1, p0+p1
		fmt.Printf("%d ", p1)
	}

	fmt.Println("]")
}

// name identifier isPrime
// takes a parameter of type int and returns a value of type bool
// last statement must be a return
func isPrime(n int) bool {
	lim := int(math.Sqrt(float64(n)))
	for p := 2; p < +lim; p++ {
		if (n % p) == 0 {
			return false
		}
	}
	return true
}

// when func id appears without parentheses, it is treated as a regular var
//   with a type and a value
//   the type of the func is determined by its signature
func add(op0 int, op1 int) int {
	return op0 + op1
}

func sub(op0, op1 int) int {
	return op0 - op1
}

// variadic parameters:
//   last parameter of a func can be delcared as variadic (variable lenght args)
//   by affixing ellipses (...) before the parameter's type
//   inidcates zero or more values of that type may be passed to the func when
//   called
func avg(nums ...float64) float64 {
	n := len(nums)
	t := 0.0
	for _, v := range nums {
		t += v
	}
	return t / float64(n)
}

func sum(nums ...float64) float64 {
	var sum float64
	for _, v := range nums {
		sum += v
	}
	return sum
}

// func result parameters
//   func can be defined to return one or more result values
//   multiple returns uses comma to separate values
func div(op0, op1 int) (int, int) {
	r := op0
	q := 0
	for r >= op1 {
		q++
		r = r - op1
	}
	// must match the return values declard in signature (2 ints)
	return q, r
}

func div2(dvdn, dvsr int) (q, r int) {
	r = dvdn
	for r >= dvsr {
		q++
		r = r - dvsr
	}
	// return statement is naked; omitting all identifiers
	// as stated in signature values for q, r will be returned to caller
	// readability preference for naked or not return
	return
}

func main() {
	// name identifier (printPi) invokes function
	printPi()

	fmt.Printf("avogadro() = %e 1/mol\n", avogadro())

	fib(41)

	prime := 37
	fmt.Printf("isPrime(%d) = %v\n", prime, isPrime(prime))

	var opAdd func(int, int) int = add
	opSub := sub
	fmt.Printf("op0(12,44)=%d\n", opAdd(12, 44))
	fmt.Printf("sub(99,13)=%d\n", opSub(99, 13))

	fmt.Printf("avg([1, 2.5, 3.75]) =%.2f\n", avg(1, 2.5, 3.75))
	points := []float64{9, 4, 3.7, 7.1, 7.9, 9.2, 10}
	// can pass a slice as a variadic parameter by adding ellipses
	fmt.Printf("sum(%v) = %.2f\n", points, sum(points...))

	q, r := div(71, 5)
	fmt.Printf("div(71,5) -> q = %d, r = %d\n", q, r)
	q2, r2 := div2(71, 5)
	fmt.Printf("div2(71,5) -> q2 = %d, r2 = %d\n", q2, r2)
}
