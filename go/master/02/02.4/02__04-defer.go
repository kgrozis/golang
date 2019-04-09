package main

import "fmt"

// defer keyword postpoines the execution of a function until the surrounding function returns
// used for file input & output operations; closing file function occurs near opening file
// also used for panic and recover builtin go functions
// deferred functions work in a lifo manner

func d1() {
	for i := 3; i > 0; i-- {
		defer fmt.Print(i, " ") // defer statement executed 3 times
	}
} // returns 1, 2, 3 (inverse order)

func d2() {
	for i := 3; i > 0; i-- {
		defer func() { // defer applied to an anonymous function
			fmt.Print(i, " ")
		}()
	} // a deferred function is evaluated after for loop ends value is 0
	fmt.Println() // returns 0, 0, 0
}

func d3() {
	for i := 3; i > 0; i-- {
		defer func(n int) {
			fmt.Print(n, " ")
		}(i) // defer applied to an anonymous function and pass value i to it
	} // each time function is deferred it uses current value of i; intentionally passing the desired value to anonymous function
} // retruns 1, 2, 3

func main() {
	d1()
	d2()
	fmt.Println()
	d3()
	fmt.Println()

}
