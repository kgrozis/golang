package main

import "fmt"

// higher-order functions
//   go can take a func as a param and return a func as a result
//   mechanism to encapsulate and abstract behaviors that can be composed
//     together to form more complex behaviors

func apply(nums []int, f func(int) int) func() {
	for i, v := range nums {
		nums[i] = f(v) // execute passed func on each element in slice
	}
	return func() { // returns func as result
		fmt.Println(nums)
	}
}

func main() {
	nums := []int{4, 32, 11, 77, 556, 3, 19, 88, 422} // slice of ints
	result := apply(nums, func(i int) int {           // inputs slice, anonymous func
		return i / 2 // anonymous func halves an element
	})
	result() // invokes returned function
}
